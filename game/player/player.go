package Player

import (
	Color "adventure_game/color"
	myLog "adventure_game/log"
	Window "adventure_game/window"
	"fmt"
)

type Player struct {
	window *Window.Window
	Pos    Window.Pos
	health int
	damage int
}

func NewPlayer(pos Window.Pos, w *Window.Window) *Player {
	return &Player{Pos: pos, window: w, health: 10, damage: 5}
}

func (p *Player) GetPos() Window.Pos {
	return p.Pos
}

func (p *Player) Attack() int {
	return p.damage
}

func (p *Player) TakeDamage(damge int) {
	p.health -= damge
	str := fmt.Sprintf("P Health: %v", p.health)
	myLog.Log(str)
}

func (p *Player) DrawUI() {
	p.window.DrawLine("Player stats", 77, 3, Color.Defalut)
	healthStr := fmt.Sprintf("Health: %v\n", p.health)
	attackStr := fmt.Sprintf("Attack: %v\n", p.Attack())
	p.window.DrawLine(healthStr, 77, 5, Color.Defalut)
	p.window.DrawLine(attackStr, 77, 7, Color.Defalut)
}
