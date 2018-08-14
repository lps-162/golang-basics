/*
	Variables declared without initial values
	are assigned zero values

	0 for numeric types
	false for boolean type
	"" for string type

*/

package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string

	fmt.Printf("%v %v %v %q \n", i, f, b, s)

}
