package main

import "fmt"

func main() {
	numList := []int{1, 2, 3}
	alphaList := []string{"a", "b", "c"}

	for _, i := range numList {
		fmt.Println(i)

		for _, letter := range alphaList {
			fmt.Println(letter)

		}
	}
	fmt.Println("\n")

	ints := [][]int{
		[]int{0, 1, 2},
		[]int{-1, -2, -3},
		[]int{9, 8, 7},
	}

	for _, i := range ints {
		fmt.Println(i)
	}
	fmt.Println("\n")

	for _, i := range ints {
		for _, j := range i {
			fmt.Println(j)
		}
	}
}
