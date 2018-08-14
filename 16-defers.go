package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("World")

	fmt.Println("Hello ")

	// Stacking defers
	fmt.Println("Stacking Defers")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("Counting Starts")
}
