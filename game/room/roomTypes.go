package Room

type RoomType int8

const (
	Empty RoomType = iota
	Start
	End
	Monster
)
