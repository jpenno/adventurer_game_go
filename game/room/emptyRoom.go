package Room

import (
	Color "adventure_game/color"
	Window "adventure_game/window"
)

type EmptyRoom struct {
	*BaseRoom
}

func NewEmptyRoom(rect Window.Rect, pos Window.Pos, w *Window.Window) EmptyRoom {
	return EmptyRoom{
		BaseRoom: NewBaseRoom(rect, pos, w, Empty, Color.Cyan, Color.Yellow, "Empty room"),
	}
}
