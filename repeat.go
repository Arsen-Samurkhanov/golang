package main

import "fmt"

func main() {
	repeat("Sunshine", 8)

	addNumbers(1, 2, 3)
}

func repeat(word string, reps int) {
	for i := 0; i < reps; i++ {
		fmt.Println(word)
	}
}

func addNumbers(x, y, z int) {
	a := x + y
	b := x + z
	c := y + z
	fmt.Println(a, b, c)
}
