package _map

import log "github.com/sirupsen/logrus"

type processor struct {
	l log.FieldLogger
}

var Processor = func(l log.FieldLogger) *processor {
	return &processor{l}
}

func (p processor) RemoveMonster(worldId byte, channelId byte, mapId uint32, id uint32) {

}

func (p processor) SpawnItem(worldId byte, channelId byte, mapId uint32, dropperId uint32, ownerId uint32, itemId uint32, x int16, y int16, quantity uint32, ffa bool, characterDrop bool) {

}
