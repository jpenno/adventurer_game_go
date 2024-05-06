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

func NewMonsterRoom(rect Window.Rect, pos Window.Pos, w Window.Window) MonsterRoom {
	tmp := MonsterRoom{
		BaseRoom: NewBaseRoom(rect, pos, w, Monster, Color.Mangenta, Color.Yellow, "Monster room"),
		enemy:    Enemy.NewEnemy(10, 1),
	}

	tmp.symble = ""
	// tmp.symble = " "

	return tmp
}

func (mr MonsterRoom) Attack(damage int) {
	if !mr.enemy.IsDead {
		mr.enemy.TakeDamage(damage)
	}

	if mr.enemy.IsDead == true {
		mr.symble = " "
	}
}

func (mr MonsterRoom) GetIsMonsterDead() bool {
	return mr.enemy.IsDead
}

func (mr MonsterRoom) GetDamage() int {
	return mr.enemy.Attack()
}

func (mr MonsterRoom) DrawUI() {

	mr.info = "Monster room"
	mr.window.DrawLine(mr.info, 77, 14, Color.Defalut)

	if mr.enemy.IsDead == false {
		monsterHealth := fmt.Sprintf("Health: %v", mr.enemy.GetHealth())
		mr.window.DrawLine(monsterHealth, 77, 16, Color.Defalut)

		monsterDamage := fmt.Sprintf("Damage: %v", mr.enemy.Attack())
		mr.window.DrawLine(monsterDamage, 77, 18, Color.Defalut)
	}
}
