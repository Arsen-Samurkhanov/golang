package main

import "fmt"

func main() {
	temperature := 40

	if temperature == 0 {
		fmt.Println("It's zero")

	} else if temperature < 0 {
		fmt.Println("It's below zero")
	} else if temperature > 0 {
		fmt.Println("It's abow zero")
	} else {
		fmt.Println("Wrong data")
	}

}
