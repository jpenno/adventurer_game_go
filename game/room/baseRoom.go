package Room

import (
	Color "adventure_game/color"
	Window "adventure_game/window"
)

type Room interface {
	Draw(active bool)
	DrawUI()
	GetType() RoomType
	GetPos() Window.Pos
}

type BaseRoom struct {
	drawRect     Window.Rect
	pos          Window.Pos
	window       *Window.Window
	roomType     RoomType
	color        Color.Color
	activeColor  Color.Color
	symble       string
	activeSymble string
	info         string
	infoPos      Window.Pos
}

func NewBaseRoom(rect Window.Rect, pos Window.Pos, window *Window.Window, roomType RoomType, color Color.Color, activeColor Color.Color, info string) *BaseRoom {
	return &BaseRoom{
		drawRect:     rect,
		window:       window,
		pos:          pos,
		roomType:     roomType,
		color:        color,
		activeColor:  activeColor,
		symble:       " ",
		activeSymble: "x",
		info:         info,
		infoPos:      Window.Pos{X: 77, Y: 3},
	}
}

func (r BaseRoom) GetType() RoomType {
	return r.roomType
}

func (r BaseRoom) GetPos() Window.Pos {
	return r.pos
}

func (r BaseRoom) Draw(active bool) {
	if active {
		r.window.DrawLine(r.activeSymble, r.drawRect.X+2, r.drawRect.Y+2, r.activeColor)
		r.window.DrawBorder(r.drawRect, r.activeColor)
	} else {
		r.window.DrawLine(r.symble, r.drawRect.X+2, r.drawRect.Y+2, r.color)
		r.window.DrawBorder(r.drawRect, r.color)
	}
}

func (r BaseRoom) DrawUI() {
	r.window.DrawLine("Room Info", r.infoPos.X, r.infoPos.Y, Color.Defalut)
	r.window.DrawLine(r.info, r.infoPos.X, r.infoPos.Y+2, Color.Defalut)
}
