package die

import (
	"math/rand"
	"strconv"
	"time"
)

// Roller rolls a n sided die
func Roller(dieAmount int, dieSides int) (int, string) {
	var output string
	if validDie(dieSides) {
		if dieAmount == 1 {
			return rollSingleDie(dieSides)
		}
		roll, output, _ := rollMultipleDie(dieAmount, dieSides)
		return roll, output
	}
	output = "Invalid Die of sides " + strconv.Itoa(dieSides)
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

func rollSingleDie(sides int) (int, string) {
	roll := rollDie(sides)
	output := "Rolled 1 " + strconv.Itoa(sides) + " sided die and got " + strconv.Itoa(roll)

	return roll, output

}

func rollMultipleDie(dieAmount int, dieSides int) (int, string, []int) {
	var finalRoll int
	toRoll := dieAmount
	var rolls []int
	for toRoll > 0 {
		roll := rollDie(dieSides)
		rolls = append(rolls, roll)
		finalRoll += roll
		toRoll--
	}
	output := "Rolled " + strconv.Itoa(dieAmount) + " " + strconv.Itoa(dieSides) + " sided die and got " + strconv.Itoa(finalRoll) + makeRollString(rolls)
	return finalRoll, output, rolls
}

func rollDie(sides int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(sides) + 1
}

func makeRollString(rolls []int) string {
	var output string
	for _, roll := range rolls {
		output += strconv.Itoa(roll) + " "
	}
	return " ( " + output + ")"
}
