package consumers

import (
	"atlas-cdc/damage"
	"github.com/sirupsen/logrus"
)

type damageCharacterCommand struct {
	CharacterId     uint32 `json:"characterId"`
	MonsterId       uint32 `json:"monsterId"`
	MonsterUniqueId uint32 `json:"monsterUniqueId"`
	DamageFrom      int8   `json:"damageFrom"`
	Element         byte   `json:"element"`
	Damage          int32  `json:"damage"`
	Direction       int8   `json:"direction"`
}

func DamageCharacterCommandCreator() EmptyEventCreator {
	return func() interface{} {
		return &damageCharacterCommand{}
	}
}

func HandleDamageCharacterCommand() EventProcessor {
	return func(l logrus.FieldLogger, e interface{}) {
		if event, ok := e.(*damageCharacterCommand); ok {
			damage.Processor(l).Damage(event.CharacterId, event.MonsterId, event.MonsterUniqueId, event.Damage, event.DamageFrom, event.Element, event.Direction)
		} else {
			l.Errorf("Unable to cast event provided to handler.")
		}
	}
}
