package main

import (
	"fmt"

	"./die"
)

func main() {

	for i := 0; i < 100; i++ {
		fmt.Println(i+1, "You rolled:", die.Roller(20))
	}
}
