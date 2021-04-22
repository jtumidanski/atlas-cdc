package character

type Model struct {
	id    uint32
	jobId uint16
	x     int16
	y     int16
	mp    uint16
	meso  uint32
}

func (m Model) Id() uint32 {
	return m.id
}

func (m Model) X() int16 {
	return m.x
}

func (m Model) Y() int16 {
	return m.y
}

func (m Model) JobId() uint16 {
	return m.jobId
}

func (m Model) MP() uint16 {
	return m.mp
}

func (m Model) Meso() uint32 {
	return m.meso
}
