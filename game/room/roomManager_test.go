package Room

import (
	Window "adventure_game/window"
	"testing"
)

func TestMakeRoomList(t *testing.T) {
	r := NewRoomManager(Window.NewWindow())

	cases := []struct {
		Name              string
		RoomTypeCount     map[RoomType]int
		WantRoomTypeCount map[RoomType]int
		RoomCount         int
	}{
		{
			Name: "check room count",
			RoomTypeCount: map[RoomType]int{
				Start:   1,
				End:     1,
				Monster: 2,
				Loot:    2,
			},
			WantRoomTypeCount: map[RoomType]int{
				Start:   1,
				End:     1,
				Monster: 2,
				Loot:    2,
			},
			RoomCount: 10,
		},
	}

	for _, test := range cases {
		got := r.makeRoomList(test.RoomTypeCount, test.RoomCount)
		gotRoomCountList := getRoomCountList(got)

		for roomType := range test.WantRoomTypeCount {
			if test.WantRoomTypeCount[roomType] != gotRoomCountList[roomType] {
				t.Errorf("%q rooms got: %d want: %d", getRoomName(roomType), gotRoomCountList[roomType], test.WantRoomTypeCount[roomType])
			}
		}
	}
}

func getRoomCountList(roomCountList []RoomType) map[RoomType]int {
	gotRoomCountList := make(map[RoomType]int)

	for _, roomType := range roomCountList {
		switch roomType {
		case Monster:
			gotRoomCountList[Monster]++
		case Loot:
			gotRoomCountList[Loot]++
		case Start:
			gotRoomCountList[Start]++
		case End:
			gotRoomCountList[End]++
		}
	}

	return gotRoomCountList
}

func getRoomName(roomType RoomType) string {

	switch roomType {
	case Monster:
		return "Monster"
	case Loot:
		return "Loot"
	case Start:
		return "Start"
	case End:
		return "End"
	}
	return "no room"
}
