package main

import "fmt"

func main() {

	fmt.Println("Demonstrating buffered channels")

	c := make(chan int, 2)
	c <- 1
	c <- 2
	c3 := func() { c <- 3 }
	go c3()

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)

}
