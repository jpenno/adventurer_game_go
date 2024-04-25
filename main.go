package main

import (
	"bufio"
	"fmt"
	"os"
)

func input() bool {
	// setPos(1, 5)
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

	window := newWindow()
	window.hideCursor()

	for {
		window.draw()
		window.printPos("x", 10, 10, Red)
		if input() {
			break
		}
	}
	window.showCursor()
}
