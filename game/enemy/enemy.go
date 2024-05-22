package Enemy

import "math/rand"

type Enemy struct {
	health uint32
	damage uint32
	IsDead bool
}

func NewEnemy(health, damage, level int) *Enemy {
	return &Enemy{
		health: uint32(rand.Intn(health)) + uint32(level),
		damage: uint32(rand.Intn(damage)) + uint32(level),
	}
}

func (e *Enemy) Attack() uint32 {
	return e.damage
}

func (e *Enemy) TakeDamage(damge uint32) {
	if e.IsDead {
		return
	}

	if e.health < damge {
		e.health = 0
		e.IsDead = true
	} else {
		e.health -= damge
		if e.health == 0 {
			e.IsDead = true
		}
	}
}

func (e *Enemy) GetHealth() uint32 {
	return e.health
}
