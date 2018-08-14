package main

import (
	"fmt"
)

// slices are dynamic flexible view of an array
// we can extract a subset of an array using slices

func main() {
	primes := []int{2, 3, 5, 7, 11, 13}

	subset := primes[1:4]

	fmt.Println("Just a subset of primes : ", subset)

	//slices are references to arrays, so if we change a
	// slice, other slices might have effect and also the original array
	clouds := [4]string{"aws", "azure", "Google Cloud", "bluemix"}

	topThree := clouds[0:3]
	promising := clouds[2:4]

	promising[0] = "gcloud"
	fmt.Println("Top Three : ", topThree)
	fmt.Println("promising : ", promising)
	fmt.Println("All clouds : ", clouds)

	// As we observed in the output, every reference to "Google Cloud"
	// has been changed to "gcloud", in slices as well as the original array

	// The following are the examples of slice literals
	// when u dont give the length in the array notation, internally an array
	// is created and then a slice is also created referring to that array
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)

	// slice defaults, 0 for the low bound, length of the slice for the high bound
	mostPopular := clouds[:2]
	mostPromising := clouds[2:]
	allClouds := clouds[:]

	fmt.Println("Demonstrating slices defaults")
	fmt.Println("Most Popular ", mostPopular)
	fmt.Println("Most Promising ", mostPromising)
	fmt.Println("All Clouds ", allClouds)
}
