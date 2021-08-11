package monster

import (
	"atlas-cdc/rest/attributes"
	log "github.com/sirupsen/logrus"
	"strconv"
)

type processor struct {
	l log.FieldLogger
}

var Processor = func(l log.FieldLogger) *processor {
	return &processor{l}
}

func (p processor) GetById(id uint32) (*Model, error) {
	resp, err := requestById(p.l)(id)
	if err != nil {
		p.l.WithError(err).Errorf("Retrieving monster %d information.", id)
		return nil, err
	}
	return makeMonster(resp.Data()), nil
}

func makeMonster(data *attributes.MonsterData) *Model {
	mid, err := strconv.ParseUint(data.Id, 10, 32)
	if err != nil {
		return nil
	}

	attr := data.Attributes
	return &Model{
		id:        uint32(mid),
		worldId:   attr.WorldId,
		channelId: attr.ChannelId,
		mapId:     attr.MapId,
		monsterId: attr.MonsterId,
	}
}

func (p processor) IsBuffed(id uint32, buff string) bool {
	p.l.Warnf("Calling into unimplemented function.")
	return false
}

func (p processor) IsNeutralized(id uint32) bool {
	return p.IsBuffed(id, "MonsterStatus.NEUTRALISE")
}

func (p processor) GetLoseItemList(monsterId uint32) ([]*LoseItem, error) {
	return make([]*LoseItem, 0), nil
}

func (p processor) DamageMonster(m *Model, characterId uint32, damage int32) {

}
