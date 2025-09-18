package main

import (
	"fmt"
)

func main() {
	flavors := []string{"chocolate", "vanila", "strawberry", "banana"}
	for _, flav := range flavors {
		if flav == "strawberry" {
			fmt.Println(flav, "is my favarite")
			continue
		}
		if flav == "vanila" {
			fmt.Println(flav, "is great")
			continue
		}

		if flav == "chocolate" {
			fmt.Println(flav, "is great!")
			continue
		}

		fmt.Println(" I've never tried", flav, "before")

	}
}
