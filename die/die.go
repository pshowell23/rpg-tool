package die

import (
	"math/rand"
	"time"
)

// Roller rolls a n sided die
func Roller(n int) int {
	if validDie(n) {
		return rollDie(n)
	}
	return 0
}

func validDie(die int) bool {
	switch die {
	case 2, 3, 4, 6, 8, 10, 12, 20, 100:
		return true
	default:
		return false
	}
}

func rollDie(sides int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(sides) + 1
}
