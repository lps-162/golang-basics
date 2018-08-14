package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4

	var f float64 = math.Sqrt(float64(x*x + y*y))

	fmt.Println("Square root of x*x + y*y is ", f)

	var z = uint(x)
	fmt.Println("z value is ", z)
}
