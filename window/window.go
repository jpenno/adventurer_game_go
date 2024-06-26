package Window

import (
	Color "adventure_game/color"
	"fmt"
)

type Rect struct {
	X      int
	Y      int
	Width  int
	Height int
}

type Pos struct {
	X int
	Y int
}

type Window struct{}

func NewWindow() *Window {
	w := Window{}

	return &w
}

func (w Window) Reset() {
	w.Clear()
	w.SetPos(0, 0)
	w.ShowCursor()
}

func (w Window) PrintPos(content string, x, y int, color Color.Color) {
	w.SetPos(x, y)
	fmt.Printf("%v", color)
	fmt.Printf(content)
}

func (w Window) SetPos(x, y int) {
	fmt.Printf("\033[%d;%dH", y, x) // Set cursor position
}

func (w Window) Clear() {
	fmt.Print("\033[2J") //Clear screen
}

func (w Window) ShowCursor() {
	fmt.Print("\x1b[?25h") // Hide cursor
}

func (w Window) HideCursor() {
	fmt.Print("\x1b[?25l") // Hide cursor
}

func (w Window) DrawBorder(rect Rect, color Color.Color) {
	char := "󰝤"
	// Draw top and bottom border
	for x := rect.X; x <= rect.X+rect.Width-1; x++ {
		w.PrintPos(char, x, rect.Y, color)
		w.PrintPos(char, x, rect.Y+rect.Height-1, color)
	}

	// Draw left and right border
	for y := rect.Y; y <= rect.Y+rect.Height-1; y++ {
		w.PrintPos(char, rect.X, y, color)
		w.PrintPos(char, rect.X+rect.Width-1, y, color)
	}
}

func (w Window) DrawLine(str string, x int, y int, color Color.Color) {
	w.PrintPos(str, x, y, color)
}
