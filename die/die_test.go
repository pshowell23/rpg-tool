package die

import (
	"testing"
)

func TestValidDieWithValidDie(t *testing.T) {
	dieSides := 20

	if !validDie(dieSides) {
		t.Fatal("Not valid die")
	}
}

func TestValidDieWithInvalidDie(t *testing.T) {
	dieSides := 5

	if validDie(dieSides) {
		t.Fatal("Valid Die")
	}
}

//TestRoller tests Roller function
func TestRollerWithValidDie(t *testing.T) {
	dieSides := 20

	roll, _ := Roller(1, dieSides)

	if roll < 1 {
		t.Fatal("Invalid Die")
	}
}

func TestRollerWithInvalidDie(t *testing.T) {
	dieSides := 5

	roll, _ := Roller(1, dieSides)

	if roll != 0 {
		t.Fatal("Valid Die")
	}
}

func TestRollerWithMultipleDie(t *testing.T) {
	dieAmount := 3
	dieSides := 20

	_, _, rolls := rollMultipleDie(dieAmount, dieSides)

	if len(rolls) != dieAmount {
		t.Fatal("Not enough dice were rolled")
	}
}

func TestMakeRollString(t *testing.T) {
	rolls := []int{1, 4, 3, 5}

	rollString := makeRollString(rolls)

	if rollString != " ( 1 4 3 5 )" {
		t.Fatal("Incorrect String")
	}
}
