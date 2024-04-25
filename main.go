package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

type colorRegistry struct {
	Red      string
	Green    string
	Yellow   string
	Blue     string
	Mangenta string
	Cyan     string
	White    string
	Defalut  string
}

func newColorRegistry() *colorRegistry {
	return &colorRegistry{
		Red:      "\033[1;31m",
		Green:    "\033[1;32m",
		Yellow:   "\033[1;33m",
		Blue:     "\033[1;34m",
		Mangenta: "\033[1;35m",
		Cyan:     "\033[1;36m",
		White:    "\033[1;37m",
		Defalut:  "\033[1;39m",
	}
}

func clear() {
	fmt.Print("\033[2J") //Clear screen
}

func setPos(x int, y int) {
	fmt.Printf("\033[%d;%dH", y, x) // Set cursor position
}

func printPos(c string, x int, y int) {
	setPos(x, y)
	color := newColorRegistry()
	r := rand.Intn(8)
	switch r {
	case 1:
		fmt.Printf(color.Green) // Set cursor position
	case 2:
		fmt.Printf(color.Red) // Set cursor position
	case 3:
		fmt.Printf(color.Yellow) // Set cursor position
	case 4:
		fmt.Printf(color.Blue) // Set cursor position
	case 5:
		fmt.Printf(color.Mangenta) // Set cursor position
	case 6:
		fmt.Printf(color.Cyan) // Set cursor position
	case 7:
		fmt.Printf(color.White) // Set cursor position
	}
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
			printPos(c, x, y)
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
