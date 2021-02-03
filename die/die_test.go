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

	roll, _ := Roller(dieSides)

	if roll < 1 {
		t.Fatal("Invalid Die")
	}
}

func TestRollerWithInvalidDie(t *testing.T) {
	dieSides := 5

	roll, _ := Roller(dieSides)

	if roll != 0 {
		t.Fatal("Valid Die")
	}
}
