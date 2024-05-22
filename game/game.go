package Game

import (
	"adventure_game/color"
	Player "adventure_game/game/player"
	Room "adventure_game/game/room"
	Window "adventure_game/window"
	"fmt"
)

type Game struct {
	window      *Window.Window
	width       int
	height      int
	roomManager *Room.RoomManager
	player      *Player.Player
	state       GameState
	controlls   *Player.Controlls
	floorLevel  uint32
}

func NewGame(window *Window.Window) *Game {
	g := Game{window: window, width: 120, height: 32}

	g.roomManager = Room.NewRoomManager(window)
	g.player = Player.NewPlayer(Window.Pos{X: 0, Y: 0}, window)
	g.player.Pos = g.roomManager.MovePlayer(g.player.Pos, g.roomManager.GetStartRoom(), g.player)
	g.state = Run
	g.controlls = Player.NewControlls()
	g.floorLevel = 1

	return &g
}

func (g *Game) Run() GameState {
	g.update()
	g.draw()
	g.input()

	return g.state
}

func (g *Game) update() {
	// reset the map ard put the player in the start room
	if g.roomManager.GetRoom(g.player.Pos).GetType() == Room.End {
		g.floorLevel++
		g.roomManager.Reset(g.floorLevel)
		g.player.Pos = g.roomManager.MovePlayer(g.player.Pos, g.roomManager.GetStartRoom(), g.player)
	}
}

func (g *Game) input() {
	g.window.DrawLine("Input: ", 3, 30, Color.Defalut)

	var input string
	fmt.Scanln(&input)

	mpos := g.player.Pos
	switch input {
	case g.controlls.Up:
		mpos.Y--
	case g.controlls.Down:
		mpos.Y++
	case g.controlls.Left:
		mpos.X--
	case g.controlls.Right:
		mpos.X++
	case g.controlls.Quit:
		g.state = Quit
	case g.controlls.Attack:
		attack(g.player, g.roomManager.GetRoom(g.player.Pos))
	case g.controlls.Pickup:
		pickup(g.player, g.roomManager.GetRoom(g.player.Pos))
	}

	g.player.Pos = g.roomManager.MovePlayer(g.player.Pos, mpos, g.player)

	if g.player.IsDead {
		g.state = PlayerDead
	}
}

func attack(player *Player.Player, room Room.Room) {
	if room.GetType() != Room.Monster {
		return
	}

	if room.(*Room.MonsterRoom).GetIsMonsterDead() {
		return
	}

	room.(*Room.MonsterRoom).Attack(player.Attack())

	if !room.(*Room.MonsterRoom).GetIsMonsterDead() {
		player.TakeDamage(room.(*Room.MonsterRoom).GetDamage())
	} else {
		player.GainXp(1)
	}
}

func pickup(player *Player.Player, room Room.Room) {
	if room.GetType() == Room.Loot {
		pickup := room.(*Room.LootRoom).Pickup()
		player.Pickup(pickup)
	}
}

func (g *Game) draw() {
	mapWidth := 5*14 + 2
	mapHeight := 5*5 + 2
	gameinfoHeight := 10
	g.window.Clear()

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

	g.player.DrawUI()
	g.roomManager.GetRoom(g.player.GetPos()).DrawUI()
}
