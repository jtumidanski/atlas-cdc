package monster

import (
	"atlas-cdc/rest/attributes"
	"atlas-cdc/rest/requests"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

const (
	monsterRegistryServicePrefix string = "/ms/morg/"
	monsterRegistryService              = requests.BaseRequest + monsterRegistryServicePrefix
	mapMonstersResource                 = monsterRegistryService + "worlds/%d/channels/%d/maps/%d/monsters"
	monstersResource                    = monsterRegistryService + "monsters"
	monsterResource                     = monstersResource + "/%d"
)

func requestInMap(l logrus.FieldLogger, span opentracing.Span) func(worldId byte, channelId byte, mapId uint32) (*attributes.MonsterDataContainer, error) {
	return func(worldId byte, channelId byte, mapId uint32) (*attributes.MonsterDataContainer, error) {
		ar := &attributes.MonsterDataContainer{}
		err := requests.Get(l, span)(fmt.Sprintf(mapMonstersResource, worldId, channelId, mapId), ar)
		if err != nil {
			return nil, err
		}
		return ar, nil
	}
}

func requestById(l logrus.FieldLogger, span opentracing.Span) func(id uint32) (*attributes.MonsterDataContainer, error) {
	return func(id uint32) (*attributes.MonsterDataContainer, error) {
		ar := &attributes.MonsterDataContainer{}
		err := requests.Get(l, span)(fmt.Sprintf(monsterResource, id), ar)
		if err != nil {
			return nil, err
		}
		return ar, nil
	}
}