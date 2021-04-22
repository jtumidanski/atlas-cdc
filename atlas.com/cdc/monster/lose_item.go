package monster

type LoseItem struct {
	itemId uint32
	x      int16
	chance int32
}

func (i LoseItem) X() int16 {
	return i.x
}

func (i LoseItem) Chance() int32 {
	return i.chance
}

func (i LoseItem) ItemId() uint32 {
	return i.itemId
}
