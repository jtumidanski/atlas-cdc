package character

import (
	"atlas-cdc/rest/attributes"
	"atlas-cdc/rest/requests"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

const (
	charactersServicePrefix     string = "/ms/cos/"
	charactersService                  = requests.BaseRequest + charactersServicePrefix
	charactersResource                 = charactersService + "characters/"
	charactersById                     = charactersResource + "%d"
	charactersInventoryResource        = charactersResource + "%d/inventories/"
	characterItems                     = charactersInventoryResource + "?type=%s&include=inventoryItems,equipmentStatistics"
	characterWeaponDamage              = charactersResource + "%d/damage/weapon"
)

func requestById(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32) (*attributes.CharacterAttributesDataContainer, error) {
	return func(characterId uint32) (*attributes.CharacterAttributesDataContainer, error) {
		ar := &attributes.CharacterAttributesDataContainer{}
		err := requests.Get(l, span)(fmt.Sprintf(charactersById, characterId), ar)
		if err != nil {
			return nil, err
		}
		return ar, nil
	}
}
