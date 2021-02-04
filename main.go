package main

import (
	"fmt"

	"./die"
)

func main() {
	fmt.Println("How many sides to roll?")
	var dieSides int
	fmt.Scan(&dieSides)
	roll, output := die.Roller(dieSides)
	fmt.Println(output)
	if roll == 0 {
		main()
	}
}
