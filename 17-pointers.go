package main

import (
	"fmt"
)

func main() {
	x, y := 100, 200

	fmt.Println(x, y)

	ptr := &x
	*ptr = 900

	fmt.Println(x, *ptr)

	ptr = &y         // point to j
	*ptr = *ptr / 37 // divide j through the pointer
	fmt.Println(y)   // see the new value of j

	// Go has no pointer arithmetic
	// otherwise Go's pointers work close to c/c++
}
