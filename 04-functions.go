package main

import "fmt"

func addNumbers(a int, b int) int {
	result := a + b
	return result
}

func addThreeNumbers(a, b, c int) int {
	result := a + b + c
	return result
}

func main() {
	var result int = addNumbers(454, 45)
	fmt.Println(result)

	fmt.Println("Sum of three numbers : ", addThreeNumbers(1, 3, 8))
}
