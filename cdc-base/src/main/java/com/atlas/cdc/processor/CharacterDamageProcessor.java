package com.atlas.cdc.processor;

import java.awt.*;
import java.util.ArrayList;
import java.util.List;
import java.util.Optional;

import com.atlas.cdc.event.producer.ShowCharacterDamageProducer;
import com.atlas.cdc.model.Character;
import com.atlas.cdc.model.LoseItem;
import com.atlas.cdc.model.Monster;
import com.atlas.cdc.model.Skill;
import com.atlas.cdc.model.SkillEffect;

public final class CharacterDamageProcessor {
   private CharacterDamageProcessor() {
   }

   public static void process(int characterId, int monsterId, int monsterUniqueId, int damage, byte damageFrom, byte element,
                              byte direction) {
      Character character = CharacterProcessor.getFromId(characterId).join();

      Monster attacker = null;
      if (damageFrom != -3 && damageFrom != -4) {
         attacker = MonsterProcessor.getFromUniqueId(monsterUniqueId).join();
         if (attacker.id() != monsterId) {
            attacker = null;
         }

         if (attacker != null) {
            if (MonsterProcessor.isNeutralized(monsterUniqueId)) {
               return;
            }

            if (damage > 0) {
               processMonsterTouchItemLoss(character, attacker);
            }
         } else if (damageFrom != 0 || !removeSelfDestructive(character.id(), monsterUniqueId)) {
            return;
         }
      }
      List<Integer> banishPlayers = new ArrayList<>();
      boolean deadly = false;
      int mpAttack = 0;
      if (damageFrom != -1 && damageFrom != -2 && attacker != null) {
         //         MobAttackInfo attackInfo = MobAttackInfoFactory.getMobAttackInfo(attacker, damageFrom);
         //         if (attackInfo != null) {
         //            if (attackInfo.deadlyAttack()) {
         //               mpAttack = character.mp() - 1;
         //                     boolean deadly = false; = true;
         //            }
         //            mpAttack += attackInfo.mpBurn();
         //            MobSkill mobSkill = MobSkillFactory.getMobSkill(attackInfo.diseaseSkill(), attackInfo.diseaseLevel());
         //            if (mobSkill != null && damage > 0) {
         //               MobSkillProcessor.getInstance().applyEffect(chr, attacker, mobSkill, false, banishPlayers);
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

      int fake = 0;
      if (damage == -1) {
         fake = 4020002 + (character.jobId() / 10 - 40) * 100000;
      }

      //      if (damage > 0) {
      //         chr.getAutoBanManager().resetMisses();
      //      } else {
      //         chr.getAutoBanManager().addMiss();
      //      }

      //in dojo player cannot use pot, so deadly attacks should be turned off as well
      int adjustedDamage = damage;
      //      if (is_deadly && GameConstants.isDojo(character.mapId()) && !YamlConfig.config.server.USE_DEADLY_DOJO) {
      //         adjustedDamage = 0;
      //         mpAttack = 0;
      //      }

      if (adjustedDamage > 0 && !CharacterProcessor.isHidden(character.id())) {
         if (attacker != null) {
            if (damageFrom == -1) {
               if (CharacterSkillProcessor.hasPowerGuard(character.id())) { // PG works on bosses, but only at half of
                  // the rate.
                  adjustedDamage = processPowerGuard(character, attacker, adjustedDamage);
               }
               if (CharacterSkillProcessor.hasBodyPressure(character.id())) {
                  processBodyPressure(character, attacker);
               }
            }

            //            MapleStatEffect cBarrier = chr.getBuffEffect(MapleBuffStat.COMBO_BARRIER);
            //            if (cBarrier != null) {
            //               adjustedDamage *= (cBarrier.getX() / 1000.0);
            //            }
         }

         if (damageFrom != -3 && damageFrom != -4) {
            int jobId = character.jobId();
            if (jobId < 200 && jobId % 10 == 2) {
               adjustedDamage = adjustDamageForAchilles(character, adjustedDamage);
            }

            adjustedDamage = adjustDamageForAranHighDefense(character, adjustedDamage);
         }

         if (CharacterSkillProcessor.hasMagicGuard(character.id()) && mpAttack == 0) {
            int magicGuard = CharacterSkillProcessor.buffValue(character.id(), "MapleBuffStat.MAGIC_GUARD");
            int mpLoss = (int) (adjustedDamage * (((double) magicGuard) / 100.0));
            int hpLoss = adjustedDamage - mpLoss;

            int currentMp = character.mp();
            if (mpLoss > currentMp) {
               hpLoss += mpLoss - currentMp;
               mpLoss = currentMp;
            }
            CharacterProcessor.adjustHealth(character.id(), -hpLoss);
            CharacterProcessor.adjustMana(character.id(), -mpLoss);
         } else if (CharacterSkillProcessor.hasMesoGuard(character.id())) {
            int mesoGuard = CharacterSkillProcessor.buffValue(character.id(), "MapleBuffStat.MESO_GUARD");
            adjustedDamage = Math.round(adjustedDamage / 2f);
            int mesoLoss = (int) (adjustedDamage * (((double) mesoGuard) / 100.0));
            if (character.meso() < mesoLoss) {
               CharacterProcessor.adjustMeso(character.id(), -character.meso(), false);
               CharacterSkillProcessor.cancelBuff(character.id(), "MapleBuffStat.MESO_GUARD");
            } else {
               CharacterProcessor.adjustMeso(character.id(), -mesoLoss, false);
            }
            CharacterProcessor.adjustHealth(character.id(), -adjustedDamage);
            CharacterProcessor.adjustMana(character.id(), -mpAttack);
         } else {
            if (CharacterProcessor.isRidingBattleship(character.id())) {
               CharacterProcessor.adjustBattleshipHp(character.id(), adjustedDamage);
            }
            CharacterProcessor.adjustHealth(character.id(), -adjustedDamage);
            CharacterProcessor.adjustMana(character.id(), -mpAttack);
         }
      }

      emitCharacterDamaged(character, attacker, damageFrom, direction, adjustedDamage, fake);

      //      if (GameConstants.isDojo(map.getId())) {
      //         chr.setDojoEnergy(chr.getDojoEnergy() + YamlConfig.config.server.DOJO_ENERGY_DMG);
      //         PacketCreator.announce(client, new GetEnergy("energy", chr.getDojoEnergy()));
      //      }

      processBanishedCharacters(attacker, banishPlayers);
   }

   /**
    * Adjusts damage taken for passive defense boost skills. Think hero achilles (and more).
    *
    * @param character the character taking damage
    * @param damage    the current damage inflicted
    * @return adjust damage taken considering skill level
    */
   protected static Integer adjustDamageForAchilles(Character character, int damage) {
      return SkillProcessor.getSkillById(character.jobId() * 10000 + (character.jobId() == 112 ? 4 : 5))
            .map(skill -> {
               int skillLevel = SkillProcessor.getSkillLevel(character.id(), skill.id());
               int effect = SkillProcessor.getSkillEffect(skill.id(), skillLevel)
                     .map(SkillEffect::x)
                     .orElse(0);
               return (int) (damage * (effect / 1000.0));
            })
            .orElse(damage);
   }

   /**
    * Adjusts damage taken for the Aran High Defense skill.
    *
    * @param character the character taking damage
    * @param damage    the current damage inflicted
    * @return adjust damage taken considering skill level
    */
   protected static int adjustDamageForAranHighDefense(Character character, int damage) {
      return SkillProcessor.getSkillById(21120004)
            .map(skill -> {
               int skillLevel = SkillProcessor.getSkillLevel(character.id(), skill.id());
               if (skillLevel > 0) {
                  int effect = SkillProcessor.getSkillEffect(skill.id(), skillLevel)
                        .map(SkillEffect::x)
                        .orElse(0);
                  return (int) (damage * (effect / 1000.0));
               }
               return damage;
            })
            .orElse(damage);
   }

   private static void processBodyPressure(Character character, Monster attacker) {
      SkillProcessor.getSkillById(21101003)
            .filter(skill -> !MonsterProcessor.isNeutralized(attacker.id()))
            .filter(skill -> !attacker.boss())
            .filter(skill -> SkillProcessor.getCharacterSkillEffect(character.id(), skill.id())
                  .map(SkillEffect::makeChanceResult)
                  .orElse(false))
            .ifPresent(skill -> {
               //               attacker.applyStatus(chr,
               //                     new MonsterStatusEffect(Collections.singletonMap(MonsterStatus.NEUTRALISE, 1), skill.get(), null,
               //                           false), false, (bPressure.getDuration() / 10) * 2, false);
            });
   }

   /**
    * Processes the effects of power guard.
    *
    * @param character the character
    * @param attacker  the attacking monster
    * @param damage    the monster has done
    * @return the adjusted damage done by the monster, after applying power guard
    */
   private static int processPowerGuard(Character character, Monster attacker, int damage) {
      double buffValue = CharacterSkillProcessor.buffValue(character.id(), "MapleBuffStat.POWER_GUARD");
      int bounceDamage = (int) (damage * (buffValue / (attacker.boss() ? 200 : 100)));
      bounceDamage = Math.min(bounceDamage, attacker.maxHp() / 10);
      int adjustedDamage = damage - bounceDamage;
      damageMonster(character, attacker, bounceDamage);
      //            attacker.aggroMonsterDamage(chr, bounceDamage);
      return adjustedDamage;
   }

   private static boolean removeSelfDestructive(int characterId, int monsterUniqueId) {
      return true;
   }

   protected static void processManaReflect(Skill skill, Character character, Monster monster, int damage) {
      if (!CharacterSkillProcessor.isBuffedFrom(character.id(), skill.id(), "MapleBuffStat.MANA_REFLECTION")) {
         return;
      }

      int skillLevel = SkillProcessor.getSkillLevel(character.id(), skill.id());
      if (skillLevel <= 0) {
         return;
      }

      Optional<SkillEffect> skillEffectOptional = SkillProcessor.getSkillEffect(skill.id(), skillLevel);
      if (skillEffectOptional.isEmpty()) {
         return;
      }

      SkillEffect skillEffect = skillEffectOptional.get();

      if (skillEffect.makeChanceResult()) {
         int bounceDamage = (damage * skillEffect.x() / 100);
         if (bounceDamage > monster.maxHp() / 5) {
            bounceDamage = monster.maxHp() / 5;
         }

         damageMonster(character, monster, bounceDamage);
         showReflectEffect(character, skill.id());
      }
   }

   protected static void showReflectEffect(Character character, int skillId) {
      //      PacketCreator.announce(character.id(), new ShowOwnBuffEffect(skillId, 5));
      //      MasterBroadcaster.getInstance()
      //            .sendToAllInMap(character.mapId(), new ShowBuffEffect(character.id(), skillId, 5, (byte) 3), false, character.id());
   }

   protected static void damageMonster(Character character, Monster attacker, int bounceDamage) {
      //      map.damageMonster(character.id(), attacker, bounceDamage);
      //      MasterBroadcaster.getInstance().sendToAllInMap(map, new DamageMonster(oid, bounceDamage), true, chr);
   }

   private static void emitCharacterDamaged(Character character, Monster monster, byte damageFrom,
                                            byte direction, int damage, int fake) {
      if (monster != null) {
         ShowCharacterDamageProducer.emit(damageFrom, monster.monsterId(), character.id(), character.mapId(), damage, fake,
               direction, false, 0, true, monster.id(), 0, 0);
      }
   }

   protected static void processBanishedCharacters(Monster monster, List<Integer> characterIds) {
      characterIds.forEach(characterId -> banishCharacter(monster, characterId));
   }

   protected static void banishCharacter(Monster monster, int characterId) {
      //player.changeMapBanish(attacker.getBanish().map(), attacker.getBanish().portal(), attacker.getBanish().msg());
   }

   protected static void processMonsterTouchItemLoss(Character character, Monster monster) {
      List<LoseItem> loseItems = MonsterProcessor.getLoseItemList(monster.monsterId());
      if (loseItems.size() != 0) {
         if (!CharacterSkillProcessor.hasAura(character.id())) {
            final int playerXPosition = character.x();
            Point pos = new Point(0, character.y());
            byte d = 1;
            for (LoseItem loseItem : loseItems) {
               //               MapleInventoryType type = ItemConstants.getInventoryType(loseItem.itemId());

               int dropCount = 0;
               //               for (byte b = 0; b < loseItem.x(); b++) {
               //                  if (Randomizer.nextInt(100) < loseItem.chance()) {
               //                     dropCount += 1;
               //                  }
               //               }

               if (dropCount > 0) {
                  int qty = CharacterProcessor.getAmountToDrop(character.id(), loseItem.itemId(), dropCount);
                  //                  removeItemFromInventory(character, type, loseItem.itemId(), qty);

                  if (loseItem.itemId() == 4031868) {
                     CharacterProcessor.updateAriantScore(character.id());
                  }

                  for (byte b = 0; b < qty; b++) {
                     pos.x = playerXPosition + ((d % 2 == 0) ? (25 * (d + 1) / 2) : -(25 * (d / 2)));
                     MapProcessor.spawnItem(character.mapId(), character.id(), loseItem.itemId());
                     d++;
                  }
               }
            }
         }
         MapProcessor.removeMonster(character.mapId(), monster.id());
      }
   }
}
