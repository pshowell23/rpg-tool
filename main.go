package main

import (
	"fmt"
	"strconv"
	"strings"

	"./die"
)

func main() {
	var input string
	fmt.Println("What do you want to roll? (eg 2d6)")
	fmt.Scan(&input)
	splitInput := strings.Split(input, "d")
	dieAmount, _ := strconv.Atoi(splitInput[0])
	dieSides, _ := strconv.Atoi(splitInput[1])
	var output string
	if splitInput[0] != "" {
		_, output = die.Roller(dieAmount, dieSides)
	} else {
		_, output = die.Roller(1, dieSides)
	}
	fmt.Println(output)
	fmt.Println("Continue? (y/n)")
	fmt.Scan(&input)
	if input == "y" {
		main()
	}
}
