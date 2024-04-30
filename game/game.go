package Game

import (
	"adventure_game/color"
	Player "adventure_game/game/player"
	Room "adventure_game/game/room"
	Window "adventure_game/window"
	"fmt"
)

type Game struct {
	window      Window.Window
	width       int
	height      int
	roomManager *Room.RoomManager
	player      *Player.Player
}

func NewGame(window Window.Window) Game {
	g := Game{window: window, width: 120, height: 32}

	g.roomManager = Room.NewRoomManager(window)
	g.player = Player.NewPlayer(Window.Pos{X: 0, Y: 0})
	g.player.Pos = g.roomManager.MovePlayer(g.player.Pos, g.roomManager.GetStartRoom())

	return g
}

func (g Game) Run() bool {
	g.update()
	g.draw()
	if !g.input() {
		return false
	}
	return true
}

func (g Game) update() {
	// reset the map ard put the player in the start room
	if g.roomManager.GetRoom(g.player.Pos.X, g.player.Pos.Y).GetType() == Room.End {
		g.roomManager.Reset()
		g.player.Pos = g.roomManager.MovePlayer(g.player.Pos, g.roomManager.GetStartRoom())
	}
}

func (g Game) input() bool {
	g.window.DrawLine("                   ", 3, 30, Color.Defalut)
	g.window.DrawLine("Input: ", 3, 30, Color.Defalut)

	var input string
	fmt.Scanln(&input)
	g.window.DrawLine("                   ", 3, 31, Color.Defalut)
	controlls := Player.NewControlls()

	mpos := g.player.Pos
	switch input {
	case controlls.Up:
		g.window.DrawLine("Move up", 3, 31, Color.Defalut)
		mpos.Y--
	case controlls.Down:
		g.window.DrawLine("down", 3, 31, Color.Defalut)
		mpos.Y++
	case controlls.Left:
		g.window.DrawLine("left", 3, 31, Color.Defalut)
		mpos.X--
	case controlls.Right:
		g.window.DrawLine("right", 3, 31, Color.Defalut)
		mpos.X++
	case controlls.Quit:
		return false
	default:
		g.window.DrawLine("bad input", 3, 31, Color.Defalut)
	}
	g.player.Pos = g.roomManager.MovePlayer(g.player.Pos, mpos)

	return true
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
