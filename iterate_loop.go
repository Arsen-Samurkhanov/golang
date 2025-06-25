package main

import "fmt"

func main() {
	sharks := []string{"hammerhead", "greate white", "dogfish", "frilled", "bullhead", "requirement"}

	for i := 0; i < len(sharks); i++ {
		fmt.Println(sharks[i])
	}
	fmt.Println("\n")

	for i, shark := range sharks {
		fmt.Println(i, shark)
	}
}
