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

	id := 0
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
					id,
				),
			)
			id++
		}
	}

	return &rm
}

func (rm *RoomManager) GetRoom(x int, y int) Room {
	return rm.rooms[y*rm.width+x]
}

func (rm *RoomManager) SetPlayerHere(x int, y int) {
	rm.playerPos = rm.rooms[y*rm.width+x].Id
}

func (rm *RoomManager) Draw() {
	for _, r := range rm.rooms {
		if r.Id == rm.playerPos {
			r.Draw(true)
		} else {
			r.Draw(false)
		}
	}
}
