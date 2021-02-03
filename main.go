package main

import (
	"fmt"

	"./die"
)

func main() {
	var dieSides int

	fmt.Println("How many sides to roll?")
	fmt.Scan(&dieSides)
	roll := die.Roller(dieSides)

	if roll == 0 {
		panic(fmt.Errorf("Invalid Die of sides %d", dieSides))
	}

	fmt.Println("Rolled", dieSides, "sided die and got", roll)

}
