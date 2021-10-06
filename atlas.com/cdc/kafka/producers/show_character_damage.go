package producers

import (
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type showCharacterDamageEvent struct {
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

func ShowCharacterDamage(l logrus.FieldLogger, span opentracing.Span) func(skillId int8, monsterId uint32, characterId uint32, mapId uint32, damage int32, fake uint32, direction int8, pgmr bool, pgmr1 byte, ispg bool, monsterUniqueId uint32, x int16, y int16) {
	producer := ProduceEvent(l, span, "TOPIC_SHOW_DAMAGE_CHARACTER_COMMAND")
	return func(skillId int8, monsterId uint32, characterId uint32, mapId uint32, damage int32, fake uint32, direction int8, pgmr bool, pgmr1 byte, ispg bool, monsterUniqueId uint32, x int16, y int16) {
		e := &showCharacterDamageEvent{
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
		producer(CreateKey(int(characterId)), e)

	}
}
