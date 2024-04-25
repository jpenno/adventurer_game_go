package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

type Color int8

const (
	Red Color = iota
	Green
	Yellow
	Blue
	Mangenta
	Cyan
	White
	Defalut
)

func (c Color) String() string {
	switch c {
	case Red:
		return "\033[1;31m"
	case Green:
		return "\033[1;32m"
	case Yellow:
		return "\033[1;33m"
	case Blue:
		return "\033[1;34m"
	case Mangenta:
		return "\033[1;35m"
	case Cyan:
		return "\033[1;36m"
	case White:
		return "\033[1;37m"
	case Defalut:
		return "\033[1;39m"
	}
	return "\033[1;39m"
}

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
	for y := 1; y <= 25; y++ {
		for x := 1; x <= 100; x++ {
			r := rand.Intn(10)
			c := strconv.Itoa(r)
			w.printPos(c, x, y, Defalut)
		}
	}
}
