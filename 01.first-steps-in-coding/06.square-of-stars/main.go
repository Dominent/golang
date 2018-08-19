package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		for x := 0; x < n; x++ {
			if ((i > 0) && (i < n-1)) && ((x > 0) && (x < n-1)) {
				fmt.Print(" ")
			} else {
				fmt.Print("*")
			}
		}
		fmt.Println()
	}
}
