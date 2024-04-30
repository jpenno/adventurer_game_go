package Player

import Window "adventure_game/window"

type Player struct {
	Pos Window.Pos
}

func NewPlayer(pos Window.Pos) *Player {
	return &Player{Pos: pos}
}

func (p Player) GetPos() Window.Pos {
	return p.Pos
}
