package Room

import (
	Color "adventure_game/color"
	Window "adventure_game/window"
)

type EndRoom struct {
	drawRect Window.Rect
	pos      Window.Pos
	window   Window.Window
	roomType RoomType
}

func NewEndRoom(rect Window.Rect, pos Window.Pos, w Window.Window) EndRoom {
	return EndRoom{drawRect: rect, pos: pos, window: w, roomType: End}
}

func (r EndRoom) GetType() RoomType {
	return r.roomType
}

func (r EndRoom) GetPos() Window.Pos {
	return r.pos
}

func (r EndRoom) Draw(active bool) {
	color := Color.Red

	if active {
		color = Color.Yellow
		r.window.DrawLine("x", r.drawRect.X+2, r.drawRect.Y+2, color)
	} else {
		r.window.DrawLine("󱊾", r.drawRect.X+2, r.drawRect.Y+2, color)
	}

	r.window.DrawBorder(r.drawRect, color)
}
