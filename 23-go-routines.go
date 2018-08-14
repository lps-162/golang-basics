package main

import "fmt"

func sumOfNumbers(arr []int, res chan int) {

	sum := 0
	for _, val := range arr {
		sum += val
	}

	res <- sum
}

func main() {

	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	myRes := make(chan int)
	go sumOfNumbers(a[:len(a)/2], myRes)
	go sumOfNumbers(a[len(a)/2:], myRes)

	x, y := <-myRes, <-myRes
	fmt.Println("Sum of numbers : ", x, y)
}
