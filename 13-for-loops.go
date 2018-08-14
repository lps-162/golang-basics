package main

import "fmt"

func main() {

	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}

	fmt.Println("Sum of first 10 numbers : ", sum)

	// The init and post statement are optional
	oneMoreSum := 1
	for oneMoreSum < 100 {
		oneMoreSum += oneMoreSum
	}

	fmt.Println(oneMoreSum)

	for {

	}
}
