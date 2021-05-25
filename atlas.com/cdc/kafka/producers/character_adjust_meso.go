package producers

import (
	"github.com/sirupsen/logrus"
)

type adjustMesoEvent struct {
	CharacterId uint32 `json:"characterId"`
	Amount      int32  `json:"amount"`
	Show        bool   `json:"bool"`
}

func AdjustMeso(l logrus.FieldLogger) func(characterId uint32, amount int32, show bool) {
	producer := ProduceEvent(l, "TOPIC_ADJUST_MESO")
	return func(characterId uint32, amount int32, show bool) {
		event := &adjustMesoEvent{characterId, amount, show}
		producer(CreateKey(int(characterId)), event)
	}
}
