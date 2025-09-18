package main

import "fmt"

func main() {
	values := []int{0, 1, 2, 3, 4, 5, 6}
	for _, x := range values {
		fmt.Println(x)
	}
}
