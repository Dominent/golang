package main

import (
	"fmt"
)

func main() {
	var a, b int

	fmt.Scan(&a)
	fmt.Scan(&b)

	square := (a * b) / 2

	fmt.Println(square)
}
