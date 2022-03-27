package character

import (
	"atlas-cdc/kafka"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type showDamageCommand struct {
	CharacterId     uint32 `json:"characterId"`
	MapId           uint32 `json:"mapId"`
	MonsterId       uint32 `json:"monsterId"`
	MonsterUniqueId uint32 `json:"monsterUniqueId"`
	SkillId         int8   `json:"skillId"`
	Damage          int32  `json:"damage"`
	Fake            uint32 `json:"fake"`
	Direction       int8   `json:"direction"`
	X               int16  `json:"x"`
	Y               int16  `json:"y"`
	PGMR            bool   `json:"pgmr"`
	PGMR1           byte   `json:"pgmr1"`
	PG              bool   `json:"pg"`
}

func emitShowDamageCommand(l logrus.FieldLogger, span opentracing.Span) func(skillId int8, monsterId uint32, characterId uint32, mapId uint32, damage int32, fake uint32, direction int8, pgmr bool, pgmr1 byte, ispg bool, monsterUniqueId uint32, x int16, y int16) {
	producer := kafka.ProduceEvent(l, span, "TOPIC_SHOW_DAMAGE_CHARACTER_COMMAND")
	return func(skillId int8, monsterId uint32, characterId uint32, mapId uint32, damage int32, fake uint32, direction int8, pgmr bool, pgmr1 byte, ispg bool, monsterUniqueId uint32, x int16, y int16) {
		e := &showDamageCommand{
			CharacterId:     characterId,
			MapId:           mapId,
			MonsterId:       monsterId,
			MonsterUniqueId: monsterUniqueId,
			SkillId:         skillId,
			Damage:          damage,
			Fake:            fake,
			Direction:       direction,
			X:               x,
			Y:               y,
			PGMR:            pgmr,
			PGMR1:           pgmr1,
			PG:              ispg,
		}
		producer(kafka.CreateKey(int(characterId)), e)
	}
}

type adjustHealthCommand struct {
	CharacterId uint32 `json:"characterId"`
	Amount      int16  `json:"amount"`
}

func emitHealthAdjustment(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32, amount int16) {
	producer := kafka.ProduceEvent(l, span, "TOPIC_ADJUST_HEALTH")
	return func(characterId uint32, amount int16) {
		c := &adjustHealthCommand{
			CharacterId: characterId,
			Amount:      amount,
		}
		producer(kafka.CreateKey(int(characterId)), c)
	}
}

type adjustManaCommand struct {
	CharacterId uint32 `json:"characterId"`
	Amount      int16  `json:"amount"`
}

func emitManaAdjustment(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32, amount int16) {
	producer := kafka.ProduceEvent(l, span, "TOPIC_ADJUST_MANA")
	return func(characterId uint32, amount int16) {
		c := &adjustManaCommand{
			CharacterId: characterId,
			Amount:      amount,
		}
		producer(kafka.CreateKey(int(characterId)), c)
	}
}

type adjustMesoCommand struct {
	CharacterId uint32 `json:"characterId"`
	Amount      int32  `json:"amount"`
	Show        bool   `json:"bool"`
}

func emitMesoAdjustment(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32, amount int32, show bool) {
	producer := kafka.ProduceEvent(l, span, "TOPIC_ADJUST_MESO")
	return func(characterId uint32, amount int32, show bool) {
		c := &adjustMesoCommand{
			CharacterId: characterId,
			Amount:      amount,
			Show:        show,
		}
		producer(kafka.CreateKey(int(characterId)), c)
	}
}
