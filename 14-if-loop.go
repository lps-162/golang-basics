package main

import (
	"fmt"
)

func main() {

	a := 10
	b := 90

	if a > b {
		fmt.Println("A is greater than B")
	} else {
		fmt.Println("B is greater than A")
	}

	if c := a * b; c >= 200 {
		fmt.Println("We reached double century")
	} else {
		fmt.Println("No we couldnt reache double century")
	}
}
