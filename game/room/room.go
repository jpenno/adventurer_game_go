package Room

import (
	Color "adventure_game/color"
	Window "adventure_game/window"
)

type Room interface {
	Draw(active bool)
	GetType() RoomType
	GetPos() Window.Pos
}

type EmptyRoom struct {
	drawRect Window.Rect
	pos      Window.Pos
	window   Window.Window
	roomType RoomType
}

func NewEmptyRoom(rect Window.Rect, pos Window.Pos, w Window.Window) EmptyRoom {
	return EmptyRoom{drawRect: rect, window: w, pos: pos, roomType: Empty}
}

func (r EmptyRoom) GetType() RoomType {
	return r.roomType
}

func (r EmptyRoom) GetPos() Window.Pos {
	return Window.Pos{X: r.drawRect.X, Y: r.drawRect.Y}
}

func (r EmptyRoom) Draw(active bool) {
	color := Color.Cyan

	if active {
		color = Color.Yellow
		r.window.DrawLine("x", r.drawRect.X+2, r.drawRect.Y+2, color)
	} else {
		r.window.DrawLine(" ", r.drawRect.X+2, r.drawRect.Y+2, color)
	}

	r.window.DrawBorder(r.drawRect, color)
}
