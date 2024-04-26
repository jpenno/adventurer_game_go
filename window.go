package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

type Window struct {
	width  int
	height int
}

func newWindow() *Window {
	w := Window{
		width:  5,
		height: 5,
	}
	return &w
}

func (w Window) printPos(c string, x int, y int, color Color) {
	w.setPos(x, y)
	fmt.Printf("%v", color)
	fmt.Printf(c)
}

func (w Window) setPos(x int, y int) {
	fmt.Printf("\033[%d;%dH", y, x) // Set cursor position
}

func (w Window) clear() {
	fmt.Print("\033[2J") //Clear screen
}

func (w Window) showCursor() {
	fmt.Print("\x1b[?25h") // Hide cursor
}

func (w Window) hideCursor() {
	fmt.Print("\x1b[?25l") // Hide cursor
}

func (w Window) draw() {
	for y := 1; y <= w.height; y++ {
		for x := 1; x <= w.width; x++ {
			r := rand.Intn(10)
			c := strconv.Itoa(r)
			w.printPos(c, x, y, Defalut)
		}
	}
}
