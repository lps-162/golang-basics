package main

import "fmt"

const Pi = 1.456

const (
	// left shifting by 100 times
	Big = 1 << 100
	// again right shifting 99 times
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	const CloudName = "gcloud"

	fmt.Printf("Hello %v \n", CloudName)

	const Truth = true
	fmt.Println("Go rules?", Truth)

	// Numeric Constants
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
