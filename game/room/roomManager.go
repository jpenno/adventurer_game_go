package Room

import (
	Player "adventure_game/game/player"
	Window "adventure_game/window"
	"fmt"
	"math/rand"
)

type RoomManageaErr string

func (e RoomManageaErr) Error() string {
	return string(e)
}

const (
	ErrTooManyRooms = RoomManageaErr("Too many rooms to add")
)

type RoomManager struct {
	width       int
	height      int
	roomWidth   int
	roomHeight  int
	roomOffSetX int
	roomOffSetY int
	rooms       []Room
	playerPos   int
	window      *Window.Window
	list        []int
}

func NewRoomManager(window *Window.Window) *RoomManager {
	rm := RoomManager{}
	rm.width = 14
	rm.height = 5
	rm.roomWidth = 5
	rm.roomHeight = 5
	rm.roomOffSetX = 3
	rm.roomOffSetY = 3
	rm.playerPos = -1
	rm.window = window
	rm.rooms = make([]Room, rm.width*rm.height)

	rm.Reset(1)

	return &rm
}

func (rm *RoomManager) Reset(floorLevel uint32) {
	roomTypeCount := map[RoomType]int{
		Monster: 5,
		Loot:    3,
		Start:   1,
		End:     1,
	}

	roomTypeList, err := rm.makeRoomList(roomTypeCount, len(rm.rooms))
	if err != nil {
		fmt.Printf("Err: %q", err)
	}

	i := 0
	for y := 0; y < rm.height; y++ {
		for x := 0; x < rm.width; x++ {
			rect := Window.Rect{
				X:      rm.roomWidth*x + rm.roomOffSetX,
				Y:      rm.roomHeight*y + rm.roomOffSetY,
				Width:  rm.roomWidth,
				Height: rm.roomHeight,
			}

			switch roomTypeList[i] {
			case Start:
				rm.rooms[i] = NewStartRoom(rect, Window.Pos{X: x, Y: y}, rm.window)
			case End:
				rm.rooms[i] = NewEndRoom(rect, Window.Pos{X: x, Y: y}, rm.window)
			case Empty:
				rm.rooms[i] = NewEmptyRoom(rect, Window.Pos{X: x, Y: y}, rm.window)
			case Monster:
				rm.rooms[i] = NewMonsterRoom(rect, Window.Pos{X: x, Y: y}, floorLevel, rm.window)
			case Loot:
				rm.rooms[i] = NewLootRoom(rect, Window.Pos{X: x, Y: y}, floorLevel, rm.window)
			}

			i++
		}
	}
}

func (rm *RoomManager) makeRoomList(roomTypeCount map[RoomType]int, size int) ([]RoomType, error) {
	roomTypeList := make([]RoomType, size)
	numberList := make([]int, size)

	totalRooms := 0
	for _, roomCount := range roomTypeCount {
		totalRooms += roomCount
	}

	if totalRooms > size {
		return nil, ErrTooManyRooms
	}

	for i := 0; i < len(numberList); i++ {
		numberList[i] = i
	}

	for roomType, roomCount := range roomTypeCount {
		for i := 0; i < roomCount; i++ {
			index := rand.Intn(len(numberList))
			roomTypeList[numberList[index]] = roomType
			numberList = remove(numberList, index)
		}
	}

	return roomTypeList, nil
}

func remove(slice []int, index int) []int {
	copy(slice[index:], slice[index+1:])
	return slice[:len(slice)-1]
}

func (rm *RoomManager) GetStartRoom() Window.Pos {
	for _, r := range rm.rooms {
		if r.GetType() == Start {
			return r.GetPos()
		}
	}

	return Window.Pos{X: -1, Y: -1}
}

func (rm *RoomManager) GetEndRoom() int {
	for i, r := range rm.rooms {
		if r.GetType() == End {
			return i
		}
	}

	return 0
}

func (rm *RoomManager) MovePlayer(ppos, mpos Window.Pos, player *Player.Player) Window.Pos {
	if ppos == mpos {
		return ppos
	}
	// Bounds checking
	if mpos.Y < 0 || mpos.Y >= rm.height || mpos.X < 0 || mpos.X >= rm.width {
		return ppos
	}

	rm.playerPos = mpos.Y*rm.width + mpos.X

	switch rm.GetRoom(ppos).GetType() {
	case Monster:
		if rm.GetRoom(ppos).(*MonsterRoom).GetIsMonsterDead() {
			return mpos
		}
		player.TakeDamage(rm.GetRoom(ppos).(*MonsterRoom).GetDamage())
	}

	return mpos
}

func (rm *RoomManager) GetRoom(pos Window.Pos) Room {
	return rm.rooms[pos.Y*rm.width+pos.X]
}

func (rm *RoomManager) Draw() {
	for i, r := range rm.rooms {
		if i == rm.playerPos {
			r.Draw(true)
		} else {
			r.Draw(false)
		}
	}
}
