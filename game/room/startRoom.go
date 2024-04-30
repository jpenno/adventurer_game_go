package Room

import (
	Color "adventure_game/color"
	Window "adventure_game/window"
)

type StartRoom struct {
	drawRect Window.Rect
	pos      Window.Pos
	window   Window.Window
	roomType RoomType
}

func NewStartRoom(rect Window.Rect, pos Window.Pos, w Window.Window) StartRoom {
	return StartRoom{drawRect: rect, pos: pos, window: w, roomType: Start}
}

func (r StartRoom) GetType() RoomType {
	return r.roomType
}

func (r StartRoom) GetPos() Window.Pos {
	return r.pos
}

func (r StartRoom) Draw(active bool) {
	color := Color.Green

	if active {
		color = Color.Yellow
		r.window.DrawLine("x", r.drawRect.X+2, r.drawRect.Y+2, color)
	} else {
		r.window.DrawLine(" ", r.drawRect.X+2, r.drawRect.Y+2, color)
	}

	r.window.DrawBorder(r.drawRect, color)
}
