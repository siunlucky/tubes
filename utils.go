package main

import (
	"math/rand"
	// "fmt"
)

func codeGenerator(min, max int) int {
	return min + rand.Intn(max-min)
}
