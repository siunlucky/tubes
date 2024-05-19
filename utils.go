package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/inancgumus/screen"
)

func codeGenerator(min, max int) int {
	return min + rand.Intn(max-min)
}

func loadingFailLogin() {
	fmt.Println("Please wait for 1 Minute to try again!")
	fmt.Println()
	time.Sleep(3 * time.Second)
	mainMenuCustomer()
}

func loadingAuth(message string) {
	fmt.Print(message, ".\n")
	time.Sleep(500 * time.Millisecond)
	fmt.Print(message, "..\n")
	time.Sleep(500 * time.Millisecond)
	fmt.Print(message, "...\n")
	time.Sleep(500 * time.Millisecond)
}

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
