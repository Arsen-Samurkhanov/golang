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
	fmt.Println("\n")

	for _, shark := range sharks {
		fmt.Println(shark)
	}
	fmt.Println("\n")

	integers := make([]int, 10)
	fmt.Println(integers)

	for i := range integers {
		integers[i] = i
	}
	fmt.Println(integers)
	fmt.Println("\n")

	sammy := "Sammy"
	for _, letter := range sammy {
		fmt.Printf("%c\n", letter)
	}
	fmt.Println("\n")

	sammyShark := map[string]string{"name": "Sammy", "animal": "shark", "color": "blue", "location": "ocean"}

	for key, value := range sammyShark {
		fmt.Println(key + ": " + value)
	}

}
