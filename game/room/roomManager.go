package Room

import (
	Player "adventure_game/game/player"
	Window "adventure_game/window"
	"math/rand"
)

type RoomManager struct {
	width        int
	height       int
	roomWidth    int
	roomHeight   int
	roomOffSetX  int
	roomOffSetY  int
	rooms        []Room
	playerPos    int
	StartRoom    Window.Pos
	window       Window.Window
	list         []int
	roomTypeList []RoomType
}

func NewRoomManager(window Window.Window) *RoomManager {
	rm := RoomManager{}
	rm.width = 14
	rm.height = 5
	rm.roomWidth = 5
	rm.roomHeight = 5
	rm.roomOffSetX = 3
	rm.roomOffSetY = 3
	rm.playerPos = -1
	rm.StartRoom = Window.Pos{X: 1, Y: 1}
	rm.window = window
	rm.rooms = make([]Room, rm.width*rm.height)

	rm.Reset()

	return &rm
}

func (rm *RoomManager) Reset() {
	rm.makeRoomList()

	i := 0
	for y := 0; y < rm.height; y++ {
		for x := 0; x < rm.width; x++ {
			rect := Window.Rect{
				X:      rm.roomWidth*x + rm.roomOffSetX,
				Y:      rm.roomHeight*y + rm.roomOffSetY,
				Width:  rm.roomWidth,
				Height: rm.roomHeight,
			}

			switch rm.roomTypeList[i] {
			case Start:
				rm.rooms[i] = NewStartRoom(rect, Window.Pos{X: x, Y: y}, rm.window)
			case End:
				rm.rooms[i] = NewEndRoom(rect, Window.Pos{X: x, Y: y}, rm.window)
			case Empty:
				rm.rooms[i] = NewEmptyRoom(rect, Window.Pos{X: x, Y: y}, rm.window)
			case Monster:
				rm.rooms[i] = NewMonsterRoom(rect, Window.Pos{X: x, Y: y}, rm.window)
			}

			i++
		}
	}
}

func (rm *RoomManager) makeRoomList() {
	rm.list = make([]int, len(rm.rooms))
	for i := 0; i < len(rm.list); i++ {
		rm.list[i] = i
	}

	rm.roomTypeList = make([]RoomType, len(rm.rooms))

	rm.addToRoomList(Start)
	rm.addToRoomList(End)

	for i := 0; i < 5; i++ {
		rm.addToRoomList(Monster)
	}
}

func (rm *RoomManager) addToRoomList(roomType RoomType) {
	index := rand.Intn(len(rm.list))
	rm.roomTypeList[rm.list[index]] = roomType
	// delete the index in the array
	rm.list = append(rm.list[:index], rm.list[index+1:]...)
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
	// Bounds checking
	if mpos.Y < 0 || mpos.Y >= rm.height || mpos.X < 0 || mpos.X >= rm.width {
		return ppos
	}

	rm.playerPos = mpos.Y*rm.width + mpos.X

	switch rm.GetRoom(ppos).GetType() {
	case Monster:
		if rm.GetRoom(ppos).(MonsterRoom).GetIsMonsterDead() {
			return mpos
		}
		player.TakeDamage(rm.GetRoom(ppos).(MonsterRoom).GetDamage())
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
