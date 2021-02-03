package die

import (
	"math/rand"
	"strconv"
	"time"
)

// Roller rolls a n sided die
func Roller(n int) (int, string) {
	var output string
	if validDie(n) {
		roll := rollDie(n)
		output = "Rolled " + strconv.Itoa(n) + " sided die and got " + strconv.Itoa(roll)
		return roll, output
	}
	output = "Invalid Die of sides " + strconv.Itoa(n)
	return 0, output
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
