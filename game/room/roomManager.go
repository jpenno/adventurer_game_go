package Room

import Window "adventure_game/window"

type RoomManager struct {
	width       int
	height      int
	roomWidth   int
	roomHeight  int
	roomOffSetX int
	roomOffSetY int
	rooms       []Room
}

func NewRoomManager(window Window.Window) RoomManager {
	rm := RoomManager{}
	rm.width = 14
	rm.height = 5
	rm.roomWidth = 5
	rm.roomHeight = 5
	rm.roomOffSetX = 3
	rm.roomOffSetY = 3

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

	return rm
}

func (rm RoomManager) Draw() {
	for _, r := range rm.rooms {
		r.Draw()
	}
}
