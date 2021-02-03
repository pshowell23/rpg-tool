package main

import (
	"fmt"

	"./die"
)

func main() {
	var dieSides int

	fmt.Println("How many sides to roll?")
	fmt.Scan(&dieSides)
	_, output := die.Roller(dieSides)
	fmt.Println(output)
}
