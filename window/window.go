package Window

import (
	"adventure_game/color"
	"fmt"
	"math/rand"
	"strconv"
)

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
