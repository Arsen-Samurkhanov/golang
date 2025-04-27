package main

import "fmt"

func main() {
	coral := [4]string{
		"blue coral",
		"staghorn coral",
		"pillar coral",
		"elkhorn coral",
	}

	fmt.Println(coral)

	// slice array
	fmt.Println(coral[1:])
	fmt.Println(coral[:3])
	fmt.Println(coral[1:3])

	// convert array to a slice
	coralSlice := coral[:]
	fmt.Println(coralSlice)

	// append new element to slice
	coralSlice = append(coralSlice, "black coral")
	fmt.Printf("%q\n", coralSlice)

}
