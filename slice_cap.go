package main

import "fmt"

func main() {
	numbers := []int{}
	for i := 0; i < 4; i++ {
		numbers = append(numbers, i)
	}
	fmt.Println(numbers)

	// more oficent way to populate the slice using cap and make

	numbers = make([]int, 4) //prealocate the size of a slice
	for i := 0; i < cap(numbers); i++ {
		numbers[i] = i
	}
	fmt.Println(numbers)

}
