package character

import (
	"atlas-cdc/kafka/producers"
	"atlas-cdc/rest/attributes"
	"errors"
	log "github.com/sirupsen/logrus"
	"math"
	"strconv"
)

type processor struct {
	l log.FieldLogger
}

var Processor = func(l log.FieldLogger) *processor {
	return &processor{l}
}

func (p processor) GetById(characterId uint32) (*Model, error) {
	cs, err := requestById(p.l)(characterId)
	if err != nil {
		return nil, err
	}

	ca := makeCharacterAttributes(cs.Data())
	if ca == nil {
		return nil, errors.New("unable to make character attributes")
	}
	return ca, nil
}

func makeCharacterAttributes(data *attributes.CharacterAttributesData) *Model {
	cid, err := strconv.ParseUint(data.Id, 10, 32)
	if err != nil {
		return nil
	}

	att := data.Attributes
	return &Model{
		id:    uint32(cid),
		jobId: att.JobId,
		x:     att.X,
		y:     att.Y,
		mp:    att.Mp,
		meso:  att.Meso,
	}
}

func (p processor) HasAura(characterId uint32) bool {
	return p.IsBuffed(characterId, "MapleBuffStat.AURA")
}

func (p processor) IsBuffed(characterId uint32, buff string) bool {
	return false
}

func (p processor) GetAmountToDrop(characterId uint32, itemId uint32, max uint32) uint32 {
	return uint32(math.Min(float64(p.CountItem(characterId, itemId)), float64(max)))
}

func (p processor) CountItem(characterId uint32, itemId uint32) uint32 {
	return 0
}

func (p processor) RemoveFromInventory(characterId uint32, itemId uint32, quantity uint32) {

}

func (p processor) UpdateAriantScore(characterId uint32) {

}

func (p processor) IsHidden(characterId uint32) bool {
	return false
}

func (p processor) HasPowerGuard(characterId uint32) bool {
	return p.IsBuffed(characterId, "MapleBuffStat.POWER_GUARD")
}

func (p processor) BuffValue(characterId uint32, buff string) float64 {
	return 0.0
}

func (p processor) HasBodyPressure(characterId uint32) bool {
	return p.IsBuffed(characterId, "MapleBuffStat.BODY_PRESSURE")
}

func (p processor) HasComboBarrier(characterId uint32) bool {
	return p.IsBuffed(characterId, "MapleBuffStat.COMBO_BARRIER")
}

func (p processor) HasMagicGuard(characterId uint32) bool {
	return p.IsBuffed(characterId, "MapleBuffStat.MAGIC_GUARD")
}

func (p processor) AdjustHealth(characterId uint32, amount int16) {
	producers.CharacterAdjustHealth(p.l)(characterId, amount)
}

func (p processor) AdjustMana(characterId uint32, amount int16) {
	producers.CharacterAdjustMana(p.l)(characterId, amount)
}

func (p processor) HasMesoGuard(characterId uint32) bool {
	return p.IsBuffed(characterId, "MapleBuffStat.MESO_GUARD")
}

func (p processor) AdjustMeso(characterId uint32, amount int32, show bool) {
	producers.AdjustMeso(p.l)(characterId, amount, show)
}

func (p processor) CancelBuff(characterId uint32, buff string) {

}

func (p processor) IsRidingBattleship(characterId uint32) bool {
	return false
}

func (p processor) AdjustBattleshipHP(characterId uint32, damage int32) {

}
