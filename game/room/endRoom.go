package Room

import (
	Color "adventure_game/color"
	Window "adventure_game/window"
)

type EndRoom struct {
	*BaseRoom
}

func NewEndRoom(rect Window.Rect, pos Window.Pos, w Window.Window) EndRoom {
	tmp := EndRoom{
		BaseRoom: NewBaseRoom(rect, pos, w, End, Color.Red, Color.Yellow, "End Room"),
	}

	tmp.symble = "󱊾"

	return tmp
}
