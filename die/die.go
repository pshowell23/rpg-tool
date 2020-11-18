package die

import (
	"math/rand"
	"time"
)

// Roller rolls a n sided die
func Roller(n int) int {
	var output int

	switch n {
	case 2:
		output = d2()
	case 3:
		output = d3()
	case 4:
		output = d4()
	case 6:
		output = d6()
	case 8:
		output = d8()
	case 10:
		output = d10()
	case 12:
		output = d12()
	case 20:
		output = d20()
	case 100:
		output = d100()
	}

	return output
}

func d2() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(2) + 1
}

func d3() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(3) + 1
}

func d4() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(4) + 1
}

func d6() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(6) + 1
}

func d8() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(8) + 1
}

func d10() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(10) + 1
}

func d12() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(12) + 1
}

func d20() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(20) + 1
}

func d100() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(100) + 1
}
