package main

import "fmt"

func main() {
	flavors := []string{"chocolate", "vanila", "strawberry", "banana"}

	for _, flav := range flavors {
		switch flav {
		case "strawberry":
			fmt.Println(flav, "is my favorite")
		case "vanila", "chocolate":
			fmt.Println(flav, "is greate")
		default:
			fmt.Println("I've never tried", flav, "before")
		}
	}

}
