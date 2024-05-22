package Room

import (
	Color "adventure_game/color"
	Enemy "adventure_game/game/enemy"
	Window "adventure_game/window"
	"fmt"
)

type MonsterRoom struct {
	*BaseRoom
	enemy *Enemy.Enemy
}

func NewMonsterRoom(rect Window.Rect, pos Window.Pos, floorLevel uint32, window *Window.Window) *MonsterRoom {
	tmp := MonsterRoom{
		BaseRoom: NewBaseRoom(rect, pos, window, Monster, Color.Mangenta, Color.Yellow, "Monster room"),
		enemy:    Enemy.NewEnemy(5, 5, int(floorLevel)),
	}

	tmp.symble = ""
	// tmp.symble = " "

	return &tmp
}

func (mr MonsterRoom) Attack(damage uint32) {
	if mr.enemy.IsDead == false {
		mr.enemy.TakeDamage(damage)
	}

	if mr.enemy.IsDead == true {
		mr.symble = " "
	}
}

func (mr MonsterRoom) GetIsMonsterDead() bool {
	return mr.enemy.IsDead
}

func (mr MonsterRoom) GetDamage() uint32 {
	return mr.enemy.Attack()
}

func (mr MonsterRoom) DrawUI() {

	mr.info = "Monster room"
	mr.window.DrawLine(mr.info, mr.infoPos.X, mr.infoPos.Y, Color.Defalut)

	if mr.enemy.IsDead == false {
		monsterHealth := fmt.Sprintf("Health: %v", mr.enemy.GetHealth())
		mr.window.DrawLine(monsterHealth, mr.infoPos.X, mr.infoPos.Y+2, Color.Defalut)

		monsterDamage := fmt.Sprintf("Damage: %v", mr.enemy.Attack())
		mr.window.DrawLine(monsterDamage, mr.infoPos.X, mr.infoPos.Y+4, Color.Defalut)
	}
}
