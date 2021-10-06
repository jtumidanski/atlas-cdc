package character

import (
	"atlas-cdc/kafka/producers"
	"atlas-cdc/rest/attributes"
	"errors"
	"github.com/opentracing/opentracing-go"
	log "github.com/sirupsen/logrus"
	"math"
	"strconv"
)

func GetById(l log.FieldLogger, span opentracing.Span) func(characterId uint32) (*Model, error) {
	return func(characterId uint32) (*Model, error) {
		cs, err := requestById(l, span)(characterId)
		if err != nil {
			return nil, err
		}

		ca := makeCharacterAttributes(cs.Data())
		if ca == nil {
			return nil, errors.New("unable to make character attributes")
		}
		return ca, nil
	}
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

func HasAura(characterId uint32) bool {
	return IsBuffed(characterId, "MapleBuffStat.AURA")
}

func IsBuffed(characterId uint32, buff string) bool {
	return false
}

func GetAmountToDrop(characterId uint32, itemId uint32, max uint32) uint32 {
	return uint32(math.Min(float64(CountItem(characterId, itemId)), float64(max)))
}

func CountItem(characterId uint32, itemId uint32) uint32 {
	return 0
}

func RemoveFromInventory(characterId uint32, itemId uint32, quantity uint32) {

}

func UpdateAriantScore(characterId uint32) {

}

func IsHidden(characterId uint32) bool {
	return false
}

func HasPowerGuard(characterId uint32) bool {
	return IsBuffed(characterId, "MapleBuffStat.POWER_GUARD")
}

func BuffValue(characterId uint32, buff string) float64 {
	return 0.0
}

func HasBodyPressure(characterId uint32) bool {
	return IsBuffed(characterId, "MapleBuffStat.BODY_PRESSURE")
}

func HasComboBarrier(characterId uint32) bool {
	return IsBuffed(characterId, "MapleBuffStat.COMBO_BARRIER")
}

func HasMagicGuard(characterId uint32) bool {
	return IsBuffed(characterId, "MapleBuffStat.MAGIC_GUARD")
}

func AdjustHealth(l log.FieldLogger, span opentracing.Span) func(characterId uint32, amount int16) {
	return func(characterId uint32, amount int16) {
		producers.CharacterAdjustHealth(l, span)(characterId, amount)
	}
}

func AdjustMana(l log.FieldLogger, span opentracing.Span) func(characterId uint32, amount int16) {
	return func(characterId uint32, amount int16) {
		producers.CharacterAdjustMana(l, span)(characterId, amount)
	}
}

func HasMesoGuard(characterId uint32) bool {
	return IsBuffed(characterId, "MapleBuffStat.MESO_GUARD")
}

func AdjustMeso(l log.FieldLogger, span opentracing.Span) func(characterId uint32, amount int32, show bool) {
	return func(characterId uint32, amount int32, show bool) {
		producers.AdjustMeso(l, span)(characterId, amount, show)
	}
}

func CancelBuff(characterId uint32, buff string) {

}

func IsRidingBattleship(characterId uint32) bool {
	return false
}

func AdjustBattleshipHP(characterId uint32, damage int32) {

}
