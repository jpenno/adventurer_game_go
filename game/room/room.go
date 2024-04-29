package Room

import (
	Color "adventure_game/color"
	Window "adventure_game/window"
	"strconv"
)

type Room struct {
	rect   Window.Rect
	window Window.Window
	Id     int
}

func NewRoom(rect Window.Rect, w Window.Window, id int) Room {
	return Room{rect: rect, window: w, Id: id}
}

func (r *Room) Draw(active bool) {
	color := Color.Cyan

	if active {
		color = Color.Yellow
	}

	r.window.DrawBorder(r.rect, color)
	r.window.DrawLine(strconv.Itoa(r.Id), r.rect.X+2, r.rect.Y+2, color)
}
