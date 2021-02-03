package die

import (
	"math/rand"
	"time"
)

// Roller rolls a n sided die
func Roller(n int) int {
	return rollDie(n)
}

func rollDie(sides int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(sides) + 1
}
