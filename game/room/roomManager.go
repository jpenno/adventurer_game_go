package Room

import (
	Window "adventure_game/window"
	"math/rand"
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
	StartRoom   Window.Pos
	window      Window.Window
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
	startPos := Window.Pos{
		X: rand.Intn(rm.width),
		Y: rand.Intn(rm.height),
	}
	endPos := Window.Pos{
		X: rand.Intn(rm.width),
		Y: rand.Intn(rm.height),
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
			if x == startPos.X && y == startPos.Y {
				rm.rooms[i] = NewStartRoom(rect, Window.Pos{X: x, Y: y}, rm.window)
			} else if x == endPos.X && y == endPos.Y {
				rm.rooms[i] = NewEndRoom(rect, Window.Pos{X: x, Y: y}, rm.window)
			} else {
				rm.rooms[i] = NewEmptyRoom(rect, Window.Pos{X: x, Y: y}, rm.window)
			}
			i++
		}
	}
}

func (rm *RoomManager) GetStartRoom() Window.Pos {
	for _, r := range rm.rooms {
		if r.GetType() == Start {
			return r.GetPos()
		}
	}
	return Window.Pos{X: -1, Y: -1}
}

func (rm *RoomManager) MovePlayer(ppos, mpos Window.Pos) Window.Pos {
	// Bounds checking
	if mpos.Y < 0 || mpos.Y >= rm.height || mpos.X < 0 || mpos.X >= rm.width {
		return ppos
	}

	rm.playerPos = mpos.Y*rm.width + mpos.X
	return mpos
}

func (rm *RoomManager) GetRoom(x int, y int) Room {
	return rm.rooms[y*rm.width+x]
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
