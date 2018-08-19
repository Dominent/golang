package main

import (
	"fmt"
)

func main() {
	const MaxStars = 10

	for i := 0; i < MaxStars; i++ {
		for x := 0; x <= i; x++ {
			fmt.Print("*")
		}

		fmt.Println()
	}
}
