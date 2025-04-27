package main

import "fmt"

func main() {
	seaCreatures := []string{"shark ", "cuttlefish ", "squid ", "mantis shrimp "}

	fmt.Println(seaCreatures)

	seaCreatures = append(seaCreatures, "seahorse")

	fmt.Println(seaCreatures)

	ocean := make([]string, 3, 5)
	fmt.Println(ocean)

}
