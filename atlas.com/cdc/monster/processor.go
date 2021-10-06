package monster

import (
	"atlas-cdc/rest/attributes"
	"github.com/opentracing/opentracing-go"
	log "github.com/sirupsen/logrus"
	"strconv"
)

func GetById(l log.FieldLogger, span opentracing.Span) func(id uint32) (*Model, error) {
	return func(id uint32) (*Model, error) {
		resp, err := requestById(l, span)(id)
		if err != nil {
			l.WithError(err).Errorf("Retrieving monster %d information.", id)
			return nil, err
		}
		return makeMonster(resp.Data()), nil
	}
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

func IsBuffed(l log.FieldLogger) func(id uint32, buff string) bool {
	return func(id uint32, buff string) bool {
		l.Warnf("Calling into unimplemented function.")
		return false
	}
}

func IsNeutralized(l log.FieldLogger) func(id uint32) bool {
	return func(id uint32) bool {
		return IsBuffed(l)(id, "MonsterStatus.NEUTRALISE")
	}
}

func GetLoseItemList(monsterId uint32) ([]*LoseItem, error) {
	return make([]*LoseItem, 0), nil
}

func DamageMonster(m *Model, characterId uint32, damage int32) {

}
