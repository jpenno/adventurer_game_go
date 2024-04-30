package Room

import (
	Color "adventure_game/color"
	Window "adventure_game/window"
)

type Room struct {
	rect   Window.Rect
	window Window.Window
}

func NewRoom(rect Window.Rect, w Window.Window) Room {
	return Room{rect: rect, window: w}
}

func (r *Room) Draw(active bool) {
	color := Color.Cyan

	if active {
		color = Color.Yellow
		r.window.DrawLine("x", r.rect.X+2, r.rect.Y+2, color)
	} else {
		r.window.DrawLine(" ", r.rect.X+2, r.rect.Y+2, color)
	}

	r.window.DrawBorder(r.rect, color)
}
