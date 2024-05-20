package Enemy

type Enemy struct {
	health uint32
	damage uint32
	IsDead bool
}

func NewEnemy(health, damage uint32) *Enemy {
	return &Enemy{health: health, damage: damage}
}

func (e *Enemy) Attack() uint32 {
	return e.damage
}

func (e *Enemy) TakeDamage(damge uint32) {
	e.health -= damge
	if e.health <= 0 {
		e.IsDead = true
	}
}

func (e *Enemy) GetHealth() uint32 {
	return e.health
}
