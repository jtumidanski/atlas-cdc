package requests

import (
	"atlas-cdc/rest/attributes"
	"fmt"
)

const (
	charactersServicePrefix     string = "/ms/cos/"
	charactersService                  = baseRequest + charactersServicePrefix
	charactersResource                 = charactersService + "characters/"
	charactersById                     = charactersResource + "%d"
	charactersInventoryResource        = charactersResource + "%d/inventories/"
	characterItems                     = charactersInventoryResource + "?type=%s&include=inventoryItems,equipmentStatistics"
	characterWeaponDamage              = charactersResource + "%d/damage/weapon"
)

var Character = func() *character {
	return &character{}
}

type character struct {
}

func (c *character) GetCharacterAttributesById(characterId uint32) (*attributes.CharacterAttributesDataContainer, error) {
	ar := &attributes.CharacterAttributesDataContainer{}
	err := get(fmt.Sprintf(charactersById, characterId), ar)
	if err != nil {
		return nil, err
	}
	return ar, nil
}
