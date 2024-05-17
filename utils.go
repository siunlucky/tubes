package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

func loading(message string) {
	clearTerminal()
	fmt.Print(message, ".\n")
	time.Sleep(500 * time.Millisecond)
	clearTerminal()
	fmt.Print(message, "..\n")
	time.Sleep(500 * time.Millisecond)
	clearTerminal()
	fmt.Print(message, "...\n")
	time.Sleep(500 * time.Millisecond)
	clearTerminal()
}

func clearTerminal() {
	screen.Clear()
	screen.MoveTopLeft()
}
