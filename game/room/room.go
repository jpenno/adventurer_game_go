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

func (r Room) Draw() {
	r.window.DrawBorder(r.rect, Color.Cyan)
}
