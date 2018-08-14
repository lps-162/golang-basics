package main

import "fmt"

func main() {
	var i, j int = 1, 2
	k := 3

	available, active, cloud := true, false, "GCloud"

	fmt.Println(i, j, k)
	fmt.Println(available, active, cloud)
}
