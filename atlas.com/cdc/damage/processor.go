package damage

import (
	"atlas-cdc/character"
	"atlas-cdc/kafka/producers"
	_map "atlas-cdc/map"
	"atlas-cdc/monster"
	"github.com/opentracing/opentracing-go"
	log "github.com/sirupsen/logrus"
	"math"
	"math/rand"
)

func Damage(l log.FieldLogger, span opentracing.Span) func(characterId uint32, monsterId uint32, monsterUniqueId uint32, damage int32, from int8, element byte, direction int8) {
	return func(characterId uint32, monsterId uint32, monsterUniqueId uint32, damage int32, from int8, element byte, direction int8) {
		c, err := character.GetById(l, span)(characterId)
		if err != nil {
			l.WithError(err).Errorf("Unable to locate character %d receiving damage.", characterId)
			return
		}

		var attacker *monster.Model
		if from != -3 && from != -4 {
			m, err := monster.GetById(l, span)(monsterUniqueId)
			if err != nil {
				l.WithError(err).Errorf("Unable to locate monster %d giving damage.", monsterUniqueId)
				return
			}
			attacker = m

			if attacker.MonsterId() != monsterId {
				attacker = nil
			}

			if attacker != nil {
				if monster.IsNeutralized(l)(attacker.Id()) {
					return
				}
				if damage > 0 {
					monsterTouchItemLoss(l)(c, attacker)
				}
			} else if from != 0 || !removeSelfDestructive(c.Id(), monsterUniqueId) {
				return
			}
		}

		banishCharacters := make([]uint32, 0)
		//deadly := false
		mpAttack := 0
		if from != -1 && from != -2 && attacker != nil {
			//         MobAttackInfo attackInfo = MobAttackInfoFactory.getMobAttackInfo(attacker, damageFrom);
			//         if (attackInfo != null) {
			//            if (attackInfo.deadlyAttack()) {
			//               mpAttack = character.mp() - 1;
			//                     boolean deadly = false; = true;
			//            }
			//            mpAttack += attackInfo.mpBurn();
			//            MobSkill mobSkill = MobSkillFactory.getMobSkill(attackInfo.diseaseSkill(), attackInfo.diseaseLevel());
			//            if (mobSkill != null && damage > 0) {
			//               MobSkillProcessor.getInstance().applyEffect(chr, attacker, mobSkill, false, banishCharacters);
			//            }
			//
			//            attacker.setMp(attacker.getMp() - attackInfo.mpCon());
			//            if (characterIsBuffed(character.id(), "MapleBuffStat.MANA_REFLECTION") && damage > 0 && !attacker.isBoss()) {
			//               int jobId = character.jobId();
			//               if (jobId == 212 || jobId == 222 || jobId == 232) {
			//                  int skillId = jobId * 10000 + 1002;
			//                  SkillProcessor.getSkillById(skillId)
			//                        .ifPresent(skill -> processManaReflect(skill, character, attacker, damage));
			//               }
			//            }
			//         }
		}

		fake := uint32(0)
		if damage == -1 {
			fake = 4020002 + (uint32(c.JobId())/10-40)*100000
		}

		//      if (damage > 0) {
		//         chr.getAutoBanManager().resetMisses();
		//      } else {
		//         chr.getAutoBanManager().addMiss();
		//      }

		//in dojo player cannot use pot, so deadly attacks should be turned off as well
		adjustedDamage := damage
		//      if (is_deadly && GameConstants.isDojo(character.mapId()) && !YamlConfig.config.server.USE_DEADLY_DOJO) {
		//         adjustedDamage = 0;
		//         mpAttack = 0;
		//      }

		if adjustedDamage > 0 && !character.IsHidden(characterId) {
			if attacker != nil {
				if from == -1 {
					if character.HasPowerGuard(characterId) {
						adjustedDamage = adjustDamageForPowerGuard(l)(c, attacker, adjustedDamage)
					}
					if character.HasBodyPressure(characterId) {
						processBodyPressure(c, attacker)
					}
					if character.HasComboBarrier(characterId) {
						adjustedDamage = adjustDamageForComboBarrier(c, adjustedDamage)
					}
				}
			}

			if from != -3 && from != -4 {
				if c.JobId() < 200 && c.JobId()%10 == 2 {
					adjustedDamage = adjustDamageForAchilles(c, adjustedDamage)
				}
				adjustedDamage = adjustDamageForAranHighDefense(c, adjustedDamage)
			}

			if mpAttack == 0 && character.HasMagicGuard(c.Id()) {
				bv := character.BuffValue(c.Id(), "MapleBuffStat.MAGIC_GUARD")
				mpLoss := int16(math.Floor(float64(adjustedDamage) * (bv / 100.0)))
				hpLoss := int16(adjustedDamage) - mpLoss

				if mpLoss > int16(c.MP()) {
					hpLoss += mpLoss - int16(c.MP())
					mpLoss = int16(c.MP())
				}
				character.AdjustHealth(l, span)(c.Id(), -hpLoss)
				character.AdjustMana(l, span)(c.Id(), -mpLoss)
			} else if character.HasMesoGuard(c.Id()) {
				bv := character.BuffValue(c.Id(), "MapleBuffStat.MESO_GUARD")
				adjustedDamage = int32(math.Round(float64(adjustedDamage) / 2))
				mesoLoss := int32(math.Floor(float64(adjustedDamage) * (bv / 100.0)))
				if c.Meso() < uint32(mesoLoss) {
					character.AdjustMeso(l, span)(c.Id(), -int32(c.Meso()), false)
					character.CancelBuff(c.Id(), "MapleBuffStat.MESO_GUARD")
				} else {
					character.AdjustMeso(l, span)(c.Id(), -mesoLoss, false)
				}
				character.AdjustHealth(l, span)(c.Id(), -int16(adjustedDamage))
				character.AdjustMana(l, span)(c.Id(), -int16(mpAttack))
			} else {
				if character.IsRidingBattleship(c.Id()) {
					character.AdjustBattleshipHP(c.Id(), adjustedDamage)
				}
				character.AdjustHealth(l, span)(c.Id(), -int16(adjustedDamage))
				character.AdjustMana(l, span)(c.Id(), -int16(mpAttack))
			}
		}

		if attacker != nil {
			producers.ShowCharacterDamage(l, span)(from, attacker.MonsterId(), c.Id(), attacker.MapId(), damage, fake, direction, false, 0, true, attacker.Id(), 0, 0)
		}

		//      if (GameConstants.isDojo(map.getId())) {
		//         chr.setDojoEnergy(chr.getDojoEnergy() + YamlConfig.config.server.DOJO_ENERGY_DMG);
		//         PacketCreator.announce(client, new GetEnergy("energy", chr.getDojoEnergy()));
		//      }

		processBanishedCharacters(attacker, banishCharacters)
	}
}

func monsterTouchItemLoss(l log.FieldLogger) func(c *character.Model, attacker *monster.Model) {
	return func(c *character.Model, attacker *monster.Model) {

		li, err := monster.GetLoseItemList(attacker.MonsterId())
		if err != nil {
			l.WithError(err).Errorf("Unable to retrieve items the monster causes the character to lose on hit.")
		}

		if len(li) == 0 {
			l.Debugf("Monster %d does not cause characters to lose items on hit.", attacker.MonsterId())
			return
		}

		if character.HasAura(c.Id()) {
			l.Debugf("Character %d has Aura buff, will not lose items.", c.Id())
			return
		}

		dropX := int16(0)
		dropY := c.Y()
		d := int16(1)
		for _, item := range li {
			dropCount := uint32(0)
			for b := int16(0); b < item.X(); b++ {
				if rand.Int31n(100) < item.Chance() {
					dropCount += 1
				}
			}

			if dropCount == 0 {
				continue
			}

			q := character.GetAmountToDrop(c.Id(), item.ItemId(), dropCount)

			character.RemoveFromInventory(c.Id(), item.ItemId(), q)

			if item.ItemId() == 4031868 {
				character.UpdateAriantScore(c.Id())
			}

			for x := uint32(0); x < q; x++ {
				if d%2 == 0 {
					dropX = c.X() + (25 * (d + 1) / 2)
				} else {
					dropX = c.X() + -(25 * (d / 2))
				}
				_map.SpawnItem(attacker.WorldId(), attacker.ChannelId(), attacker.MapId(), c.Id(), c.Id(), item.ItemId(), dropX, dropY, q, true, true)
				d++
			}
		}

		_map.RemoveMonster(attacker.WorldId(), attacker.ChannelId(), attacker.MapId(), attacker.Id())
	}
}

func removeSelfDestructive(characterId uint32, monsterUniqueId uint32) bool {
	return true
}

func adjustDamageForPowerGuard(l log.FieldLogger) func(c *character.Model, m *monster.Model, damage int32) int32 {
	return func(c *character.Model, m *monster.Model, damage int32) int32 {
		bv := character.BuffValue(c.Id(), "MapleBuffStat.POWER_GUARD")
		md := 100
		if m.IsBoss() {
			md = 200
		}

		bounceDamage := int32(math.Floor(float64(damage) * (bv / float64(md))))
		adjustedDamage := damage - bounceDamage

		monster.DamageMonster(m, c.Id(), bounceDamage)

		return adjustedDamage
	}
}

func processBodyPressure(c *character.Model, m *monster.Model) {
	//SkillProcessor.getSkillById(21101003)
	//.filter(skill -> !MonsterProcessor.isNeutralized(attacker.id()))
	//.filter(skill -> !attacker.boss())
	//.filter(skill -> SkillProcessor.getCharacterSkillEffect(character.id(), skill.id())
	//.map(SkillEffect::makeChanceResult)
	//.orElse(false))
	//.ifPresent(skill -> {
	//	//               attacker.applyStatus(chr,
	//	//                     new MonsterStatusEffect(Collections.singletonMap(MonsterStatus.NEUTRALISE, 1), skill.get(), null,
	//	//                           false), false, (bPressure.getDuration() / 10) * 2, false);
	//});
}

func adjustDamageForComboBarrier(c *character.Model, damage int32) int32 {
	//adjustedDamage *= (cBarrier.getX() / 1000.0);
	return damage
}

func adjustDamageForAchilles(c *character.Model, damage int32) int32 {
	//return SkillProcessor.getSkillById(character.jobId() * 10000 + (character.jobId() == 112 ? 4 : 5))
	//.map(skill -> {
	//	int skillLevel = SkillProcessor.getSkillLevel(character.id(), skill.id());
	//	int effect = SkillProcessor.getSkillEffect(skill.id(), skillLevel)
	//	.map(SkillEffect::x)
	//	.orElse(0);
	//	return (int) (damage * (effect / 1000.0));
	//})
	//.orElse(damage);
	return damage
}

func adjustDamageForAranHighDefense(c *character.Model, damage int32) int32 {
	//return SkillProcessor.getSkillById(21120004)
	//.map(skill -> {
	//	int skillLevel = SkillProcessor.getSkillLevel(character.id(), skill.id());
	//	if (skillLevel > 0) {
	//		int effect = SkillProcessor.getSkillEffect(skill.id(), skillLevel)
	//		.map(SkillEffect::x)
	//		.orElse(0);
	//		return (int) (damage * (effect / 1000.0));
	//	}
	//	return damage;
	//})
	//.orElse(damage);
	return damage
}

func processBanishedCharacters(attacker *monster.Model, characters []uint32) {
	for _, cid := range characters {
		banishCharacter(attacker, cid)
	}
}

func banishCharacter(attacker *monster.Model, characterId uint32) {
	//player.changeMapBanish(attacker.getBanish().map(), attacker.getBanish().portal(), attacker.getBanish().msg());
}
