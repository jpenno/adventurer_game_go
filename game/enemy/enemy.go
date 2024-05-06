package Enemy

type Enemy struct {
	health int
	damage int
	IsDead bool
}

func NewEnemy(h, d int) *Enemy {
	return &Enemy{health: h, damage: d}
}

func (e *Enemy) Attack() int {
	return e.damage
}

func (e *Enemy) TakeDamage(damge int) {
	e.health -= damge
	if e.health <= 0 {
		e.IsDead = true
	}
}

func (e *Enemy) GetHealth() int {
	return e.health
}
