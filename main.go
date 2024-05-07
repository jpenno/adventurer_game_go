package main

import (
	"adventure_game/game"
	"adventure_game/window"
)

func main() {
	window := Window.NewWindow()
	window.HideCursor()
	window.Clear()

	game := Game.NewGame(window)

	for {
		switch game.Run() {
		case Game.Quit:
			window.Reset()
			return
		}
	}
}
