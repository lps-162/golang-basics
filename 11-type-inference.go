package main

import (
	"fmt"
)

func main() {
	var i int = 56

	j := i // type inference

	x := 42
	y := 3.142
	cloud := "gcloud"

	fmt.Printf("Types %T %T \n", i, j)
	fmt.Printf("Types %T %T %T ", x, y, cloud)
}
