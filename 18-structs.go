package main

import "fmt"

// structs are collection of fields
// members can accessed by dots
type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{14, 26}
	fmt.Println(v)

	// printing out individual members using dot syntax
	fmt.Println("Vertex X : ", v.X, " Y : ", v.Y)

	// we can have pointers to structs
	vPtr := &v

	// Go allows direct usage without * operator for dereferncing pointers to strcuts
	// instead of (*vPtr).X, we can access vPtr.X
	vPtr.X = 189

	fmt.Println("After Modificatio thru struct pointers X : ", vPtr.X, " Y : ", vPtr.Y)

	yZeroVertex := Vertex{X: 88}
	bothZeroVertex := Vertex{}
	ptrToVertex := &Vertex{5, 6}

	// printing the new set of vertices
	fmt.Println("Printing the new set of vertices")
	fmt.Println(v, yZeroVertex, bothZeroVertex, ptrToVertex)
}
