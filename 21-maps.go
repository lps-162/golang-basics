package main

import "fmt"

// Vertex ...co-ordinates of the specific location
type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Trivandrum"] = Vertex{59.8789, -53.5634}

	fmt.Println(m["Trivandrum"])

	cities := map[string]Vertex{
		"Kochin":    Vertex{14.234, 45.9234},
		"Kothagiri": Vertex{89.999, 45.830},
	}

	fmt.Println("List of Cities")
	fmt.Println(cities)

}
