package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Println("Continue loop")
			continue
		}
		fmt.Println("The value of i is", i)
	}
	fmt.Println("Exiting program")
}
