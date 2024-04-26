package Window

import (
	Color "adventure_game/color"
	"fmt"
	"math/rand"
	"strconv"
)

type Rect struct {
	X      int
	Y      int
	Width  int
	Height int
}

type Window struct {
	width  int
	height int
}

func NewWindow() *Window {
	w := Window{
		width:  5,
		height: 5,
	}
	return &w
}

func (w Window) PrintPos(c string, x int, y int, color Color.Color) {
	w.SetPos(x, y)
	fmt.Printf("%v", color)
	fmt.Printf(c)
}

func (w Window) SetPos(x int, y int) {
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

func (w Window) Draw() {
	for y := 1; y <= w.height; y++ {
		for x := 1; x <= w.width; x++ {
			r := rand.Intn(10)
			c := strconv.Itoa(r)
			w.PrintPos(c, x, y, Color.Defalut)
		}
	}
}

func (w Window) DrawBorder(rect Rect, color Color.Color) {
	char := "󰝤"
	// Draw top and bottom border
	for x := rect.X; x <= rect.Width; x++ {
		w.PrintPos(char, x, rect.Y, color)
		w.PrintPos(char, x, rect.Y+rect.Height-1, color)
	}

	// Draw left and right border
	for y := rect.Y; y <= rect.Height; y++ {
		w.PrintPos(char, rect.X, y, color)
		w.PrintPos(char, rect.X+rect.Width-1, y, color)
	}
}
