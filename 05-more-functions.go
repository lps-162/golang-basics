package main

import "fmt"

func swapNumbers(a, b int) (int, int) {

	return b, a
}

func swapNames(firstName, lastName string) (string, string) {
	return lastName, firstName
}

func main() {
	x := 20
	y := 90

	x, y = swapNumbers(x, y)

	fmt.Println("X and Y after swapping : ", x, y)

	fname, lname := swapNames("Shivan", "LP")

	fmt.Println("Swapped names", fname, lname)
}
