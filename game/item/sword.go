package item

import (
	"math/rand"
)

type Sword struct {
	*baseItem
}

func NewSword(name string, level uint32) *Sword {
	stats := make(map[string]uint32)
	stats["level"] = level
	stats["damage"] = level + uint32(rand.Intn(4))
	return &Sword{
		baseItem: &baseItem{icon: "󰓥", name: name, stats: stats},
	}
}

func (s Sword) Use() uint32 {
	return s.stats["damage"]
}

func (s Sword) GetDamage() uint32 {
	return s.stats["damage"]
}
func (s Sword) GetLevel() uint32 {
	return s.stats["level"]
}
