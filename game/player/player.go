package Player

import (
	Color "adventure_game/color"
	Item "adventure_game/game/item"
	myLog "adventure_game/log"
	Window "adventure_game/window"
	"fmt"
)

type Player struct {
	window    *Window.Window
	Pos       Window.Pos
	sword     *Item.Sword
	health    uint32
	maxHealth uint32
	damage    uint32
	level     uint32
	xp        uint32
	xpToLevel uint32
	infoPos   Window.Pos
}

func NewPlayer(pos Window.Pos, window *Window.Window) *Player {
	return &Player{
		Pos:       pos,
		window:    window,
		health:    10,
		maxHealth: 10,
		damage:    5,
		level:     1,
		xp:        0,
		xpToLevel: 5,
		sword:     nil,
		infoPos:   Window.Pos{X: 77, Y: 14},
	}
}

func (p *Player) GetPos() Window.Pos {
	return p.Pos
}

func (p *Player) getDamage() uint32 {
	return p.damage + p.level
}

func (p *Player) Attack() uint32 {
	if p.sword != nil {
		return p.getDamage() + p.sword.Use()
	}
	return p.getDamage()
}

func (p *Player) GainXp(xpGain uint32) {
	p.xp += xpGain
	if p.xp >= p.xpToLevel {
		p.level++
		p.xpToLevel++
		p.xp = 0
	}
}

func (p *Player) TakeDamage(damge uint32) {
	p.health -= damge
	str := fmt.Sprintf("P Health: %v", p.health)
	myLog.Log(str)
}

func (p *Player) Pickup(item Item.Item) {
	p.sword = item.(*Item.Sword)
}

func (p *Player) DrawUI() {
	healthStr := fmt.Sprintf("Health: %v/%v\n", p.health, p.maxHealth)
	attackStr := fmt.Sprintf("Attack: %v\n", p.Attack())
	levelStr := fmt.Sprintf("Level: %v\n", p.level)
	xpStr := fmt.Sprintf("Xp: %v\n", p.xp)
	xpToLevelStr := fmt.Sprintf("Xp To Level: %v\n", p.xpToLevel)

	p.window.DrawLine("Player stats", p.infoPos.X, p.infoPos.Y, Color.Defalut)
	p.window.DrawLine(levelStr, p.infoPos.X, p.infoPos.Y+1, Color.Defalut)
	p.window.DrawLine(healthStr, p.infoPos.X, p.infoPos.Y+2, Color.Defalut)
	p.window.DrawLine(attackStr, p.infoPos.X, p.infoPos.Y+3, Color.Defalut)
	p.window.DrawLine(xpStr, p.infoPos.X, p.infoPos.Y+4, Color.Defalut)
	p.window.DrawLine(xpToLevelStr, p.infoPos.X, p.infoPos.Y+5, Color.Defalut)

	if p.sword != nil {
		nameStr := fmt.Sprintf("Weapon: %v\n", p.sword.GetName())
		p.window.DrawLine(nameStr, p.infoPos.X+20, p.infoPos.Y, Color.Defalut)
		attackStr := fmt.Sprintf("Attack: %v\n", p.sword.GetDamage())
		p.window.DrawLine(attackStr, p.infoPos.X+20, p.infoPos.Y+1, Color.Defalut)
	}
}
