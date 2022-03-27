package damage

import (
	"atlas-cdc/kafka/consumers"
	"atlas-cdc/kafka/handler"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

const (
	consumerName = "damage_character_command"
	topicToken   = "DAMAGE_CHARACTER"
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

func NewConsumer(groupId string) consumers.Config[damageCharacterCommand] {
	return consumers.NewConfiguration[damageCharacterCommand](consumerName, topicToken, groupId, HandleDamageCharacterCommand())
}

func HandleDamageCharacterCommand() handler.EventHandler[damageCharacterCommand] {
	return func(l logrus.FieldLogger, span opentracing.Span, command damageCharacterCommand) {
		Damage(l, span)(command.CharacterId, command.MonsterId, command.MonsterUniqueId, command.Damage, command.DamageFrom, command.Element, command.Direction)
	}
}
