package main

import (
	"adventure_game/game"
	"adventure_game/window"
	"fmt"
)

func main() {
	window := Window.NewWindow()
	window.HideCursor()
	window.Clear()

	game := Game.NewGame(window)

	for {
		switch game.Run() {
		case Game.PlayerDead:
			window.Reset()
			fmt.Print("Game over player died\n")
			return
		case Game.Quit:
			window.Reset()
			return
		}
	}
}
