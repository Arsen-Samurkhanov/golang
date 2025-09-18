package main

import "fmt"

func main() {
	permissons := map[int]string{
		1: "read",
		2: "write",
		3: "delete",
	}
	fmt.Println(permissons)

	delete(permissons, 3)
	fmt.Println(permissons)

	//remove all itmems by assigning to an empty map
	permissons = map[int]string{}
	fmt.Println(permissons)

}
