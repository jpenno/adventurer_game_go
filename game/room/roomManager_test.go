package Room

import (
	Window "adventure_game/window"
	"testing"
)

func TestMakeRoomList(t *testing.T) {
	r := NewRoomManager(Window.NewWindow())

	wantMonsterRooms := 5
	wantLootRooms := 3

	rtc := make(map[RoomType]int)
	rtc[Monster] = wantMonsterRooms
	rtc[Loot] = wantLootRooms

	got := r.makeRoomList(rtc)
	gotMonsterRooms := 0
	gotLootRooms := 0

	for _, rt := range got {
		switch rt {
		case Monster:
			gotMonsterRooms++
		case Loot:
			gotLootRooms++
		}
	}

	if wantMonsterRooms != gotMonsterRooms {
		t.Errorf("want monster roomsa: %d, got monster rooms: %d", wantMonsterRooms, gotMonsterRooms)
	}
	if wantLootRooms != gotLootRooms {
		t.Errorf("want monster roomsa: %d, got monster rooms: %d", wantLootRooms, gotLootRooms)
	}
}
