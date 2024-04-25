package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
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

func clear() {
	fmt.Print("\033[2J") //Clear screen
}

func setPos(x int, y int) {
	fmt.Printf("\033[%d;%dH", y, x) // Set cursor position
}

func printPos(c string, x int, y int, color Color) {
	setPos(x, y)
	fmt.Printf("%v", color)
	fmt.Printf(c)
}

func hideCursor() {
	fmt.Print("\x1b[?25l") // Hide cursor
}

func showCursor() {
	fmt.Print("\x1b[?25h") // Hide cursor
}

func draw() {
	for y := 1; y <= 25; y++ {
		for x := 1; x <= 100; x++ {
			r := rand.Intn(10)
			c := strconv.Itoa(r)
			printPos(c, x, y, Defalut)
		}
	}
}

func input() bool {
	setPos(1, 5)
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
	clear()
	hideCursor()

	for {
		draw()
		if input() {
			break
		}
	}
	showCursor()
}
