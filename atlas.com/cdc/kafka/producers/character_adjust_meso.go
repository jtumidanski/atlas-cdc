package producers

import (
	"context"
	log "github.com/sirupsen/logrus"
)

type adjustMesoEvent struct {
	CharacterId uint32 `json:"characterId"`
	Amount      int32  `json:"amount"`
	Show        bool   `json:"bool"`
}

var AdjustMeso = func(l log.FieldLogger, ctx context.Context) *adjustMeso {
	return &adjustMeso{
		l:   l,
		ctx: ctx,
	}
}

type adjustMeso struct {
	l   log.FieldLogger
	ctx context.Context
}

func (e *adjustMeso) Emit(characterId uint32, amount int32, show bool) {
	event := &adjustMesoEvent{characterId, amount, show}
	produceEvent(e.l, "TOPIC_ADJUST_MESO", createKey(int(characterId)), event)
}