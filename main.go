package main

import (
	"fmt"

	"./die"
)

func main() {
	var dieSides int

	fmt.Println("How many sides to roll?")
	fmt.Scan(&dieSides)

	fmt.Println("Rolled", dieSides, "sided die and got", die.Roller(dieSides))

}
