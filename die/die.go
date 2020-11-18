package die

import (
	"math/rand"
	"time"
)

// Roller rolls a n sided die
func Roller(n int) int {
	rand.Seed(time.Now().UnixNano())
	roll := rand.Intn(n) + 1

	return roll
}
