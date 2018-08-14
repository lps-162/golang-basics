package main

import (
	"fmt"
	"math/rand"
)

func main() {
	rand.Seed(1435)
	randomNumber := rand.Intn(100)
	fmt.Println("Generating a random number ", randomNumber)
}
