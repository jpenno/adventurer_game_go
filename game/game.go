package Game

import (
	"adventure_game/color"
	Room "adventure_game/game/room"
	Window "adventure_game/window"
)

type Game struct {
	window      Window.Window
	width       int
	height      int
	roomManager *Room.RoomManager
}

func NewGame(window Window.Window) Game {
	g := Game{window: window, width: 120, height: 32}

	g.roomManager = Room.NewRoomManager(window)

	return g
}

func (g Game) Run() {
	g.update()
	g.draw()
}

func (g Game) update() {
	g.roomManager.SetPlayerHere(0, 0)
}

func (g Game) draw() {
	mapWidth := 5*14 + 2
	mapHeight := 5*5 + 2
	gameinfoHeight := 10

	// Game border
	g.window.DrawBorder(Window.Rect{X: 1, Y: 1, Width: g.width, Height: g.height}, Color.Red)

	// Map border
	g.window.DrawBorder(Window.Rect{X: 2, Y: 2, Width: mapWidth, Height: mapHeight}, Color.Blue)

	// Game info
	g.window.DrawBorder(Window.Rect{X: mapWidth + 2, Y: 2, Width: g.width - mapWidth - 2, Height: gameinfoHeight}, Color.Green)

	// Inventory
	g.window.DrawBorder(Window.Rect{X: mapWidth + 2, Y: gameinfoHeight + 3, Width: g.width - mapWidth - 2, Height: 10}, Color.Green)

	// Rooms
	g.roomManager.Draw()
}
