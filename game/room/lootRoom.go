package Room

import (
	Color "adventure_game/color"
	"adventure_game/game/item"
	Window "adventure_game/window"
	"fmt"
)

type LootRoom struct {
	*BaseRoom
	item *item.Sword
}

func NewLootRoom(rect Window.Rect, pos Window.Pos, floorLevel uint32, window *Window.Window) *LootRoom {
	tmp := LootRoom{
		BaseRoom: NewBaseRoom(rect, pos, window, Loot, Color.Green, Color.Yellow, "Loot room"),
		item:     item.NewSword("sword", floorLevel),
	}

	tmp.symble = tmp.item.GetIcon()

	return &tmp
}

func (lr *LootRoom) Pickup() item.Item {
	pickup := lr.item
	lr.item = nil
	lr.symble = " "
	return pickup
}

func (lr LootRoom) DrawUI() {
	lr.info = "Loot room"
	lr.window.DrawLine(lr.info, lr.infoPos.X, lr.infoPos.Y, Color.Defalut)

	if lr.item != nil {
		itemName := fmt.Sprintf("Item: %v", lr.item.GetName())
		lr.window.DrawLine(itemName, lr.infoPos.X, lr.infoPos.Y+1, Color.Defalut)

		stats := lr.item.GetStats()
		yOffSet := 2
		for key, stat := range stats {
			statStr := fmt.Sprintf("%v: %v", key, stat)
			lr.window.DrawLine(statStr, lr.infoPos.X, lr.infoPos.Y+yOffSet, Color.Defalut)
			yOffSet++
		}
	}
}
