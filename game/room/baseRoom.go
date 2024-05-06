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
	window       Window.Window
	roomType     RoomType
	color        Color.Color
	activeColor  Color.Color
	symble       string
	activeSymble string
	info         string
}

func NewBaseRoom(rect Window.Rect, pos Window.Pos, w Window.Window, rt RoomType, color Color.Color, activeColor Color.Color, info string) *BaseRoom {
	return &BaseRoom{
		drawRect:     rect,
		window:       w,
		pos:          pos,
		roomType:     rt,
		color:        color,
		activeColor:  activeColor,
		symble:       " ",
		activeSymble: "x",
		info:         info,
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
	r.window.DrawLine("Room Info", 77, 14, Color.Defalut)
	r.window.DrawLine(r.info, 77, 16, Color.Defalut)
}
