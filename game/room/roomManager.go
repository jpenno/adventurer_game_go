package Room

import (
	Window "adventure_game/window"
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

	for y := 0; y < rm.height; y++ {
		for x := 0; x < rm.width; x++ {
			rm.rooms = append(
				rm.rooms,
				NewRoom(
					Window.Rect{
						X:      rm.roomWidth*x + rm.roomOffSetX,
						Y:      rm.roomHeight*y + rm.roomOffSetY,
						Width:  rm.roomWidth,
						Height: rm.roomHeight,
					},
					window,
				),
			)
		}
	}

	return &rm
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
