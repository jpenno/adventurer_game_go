package Player

import (
	Color "adventure_game/color"
	"adventure_game/game/item"
	myLog "adventure_game/log"
	Window "adventure_game/window"
	"fmt"
)

type Player struct {
	window  *Window.Window
	Pos     Window.Pos
	sword   *item.Sword
	health  uint32
	damage  uint32
	infoPos Window.Pos
}

func NewPlayer(pos Window.Pos, w *Window.Window) *Player {
	return &Player{
		Pos:     pos,
		window:  w,
		health:  10,
		damage:  5,
		sword:   nil,
		infoPos: Window.Pos{X: 77, Y: 14},
	}
}

func (p *Player) GetPos() Window.Pos {
	return p.Pos
}

func (p *Player) Attack() uint32 {
	if p.sword != nil {
		return p.damage + p.sword.Use()
	}
	return p.damage
}

func (p *Player) TakeDamage(damge uint32) {
	p.health -= damge
	str := fmt.Sprintf("P Health: %v", p.health)
	myLog.Log(str)
}

func (p *Player) Pickup(i item.Item) {
	p.sword = i.(*item.Sword)
}

func (p *Player) DrawUI() {
	p.window.DrawLine("Player stats", p.infoPos.X, p.infoPos.Y, Color.Defalut)
	healthStr := fmt.Sprintf("Health: %v\n", p.health)
	attackStr := fmt.Sprintf("Attack: %v\n", p.Attack())
	p.window.DrawLine(healthStr, p.infoPos.X, p.infoPos.Y+2, Color.Defalut)
	p.window.DrawLine(attackStr, p.infoPos.X, p.infoPos.Y+4, Color.Defalut)
}
