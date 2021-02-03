package main

import (
	"fmt"

	"./die"
)

func main() {
	fmt.Println("How many sides to roll?")
	var dieSides int
	fmt.Scan(&dieSides)
	_, output := die.Roller(dieSides)
	fmt.Println(output)
}
