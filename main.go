package main

import (
	"adventure_game/color"
	"adventure_game/window"
	"bufio"
	"fmt"
	"os"
)

func input() bool {
	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()

	if err != nil {
		fmt.Println(err)
	}
	if char == 'q' {
		return true
	}
	fmt.Printf("\n")
	fmt.Printf("char: %v\n", char)
	return false
}

func main() {
	fmt.Println("hello world")

	window := Window.NewWindow()
	window.HideCursor()
	window.Clear()

	for {
		window.Draw()
		window.PrintPos("x", 10, 10, Color.Red)
		if input() {
			break
		}
	}
	window.ShowCursor()
}
