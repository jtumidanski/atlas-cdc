package monster

type Model struct {
	id        uint32
	worldId   byte
	channelId byte
	mapId     uint32
	monsterId uint32
}

func (m Model) MonsterId() uint32 {
	return m.monsterId
}

func (m Model) Id() uint32 {
	return m.id
}

func (m Model) WorldId() byte {
	return m.worldId
}

func (m Model) ChannelId() byte {
	return m.channelId
}

func (m Model) MapId() uint32 {
	return m.mapId
}

func (m Model) IsBoss() bool {
	return false
}
