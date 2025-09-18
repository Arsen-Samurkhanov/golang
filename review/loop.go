package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("\n")
	for i := 0; i < 15; i += 3 {
		fmt.Println(i)
	}

	fmt.Println("\n")
	for i := 100; i > 0; i -= 10 {
		fmt.Println(i)
	}

	fmt.Println("\n")
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	fmt.Println("\n")
	i = 0
	for {
		i += 5
		fmt.Println(i)
		if i == 30 {
			break
		}
	}

}
