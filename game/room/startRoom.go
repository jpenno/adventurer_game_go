package Room

import (
	Color "adventure_game/color"
	Window "adventure_game/window"
)

type StartRoom struct {
	BaseRoom
}

func NewStartRoom(rect Window.Rect, pos Window.Pos, w Window.Window) StartRoom {
	return StartRoom{
		BaseRoom: NewBaseRoom(rect, pos, w, Start, Color.Green, Color.Yellow),
	}
	// return StartRoom{
	// 	BaseRoom{
	// 		drawRect:    rect,
	// 		pos:         pos,
	// 		window:      w,
	// 		roomType:    Start,
	// 		color:       Color.Green,
	// 		activeColor: Color.Yellow,
	// 	},
	// }
}
