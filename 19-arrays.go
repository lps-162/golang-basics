package main

import "fmt"

type Actor struct {
	Name     string
	Industry string
}

func main() {

	// Arrays length is part of its type,
	// So its size cannot be changed
	// But we have slices to work with arrays
	var cities [2]string
	cities[0] = "Cochin"
	cities[1] = "Trivandrum"

	fmt.Println(cities[0], cities[1])
	fmt.Println(cities)

	magicNumbers := [5]int{1, 5, 9, 45, 90}
	fmt.Println(magicNumbers)

	// we can have arrays of structs also
	actors := [3]Actor{}
	actors[0] = Actor{"Rock Dwayne", "Hollywood"}
	actors[1] = Actor{"Deepika Padukone", "Hollywood"}
	actors[2] = Actor{"Ranveer Singh", "Bollywood"}

	for i := 0; i < len(actors); i++ {
		fmt.Printf("Actor %v : %v \n", i, actors[i])
	}
}
