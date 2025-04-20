package main

import "fmt"

func main() {
	counts := map[string]int{}
	fmt.Println(counts["sammy"])

	count, ok := counts["sammy"]

	if ok {
		fmt.Printf("Sammy has count of %d\n", count)
	} else {
		fmt.Println("sammy was not found")
	}

}
