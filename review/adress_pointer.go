package main

import "fmt"

func main() {
	var creature string = "shark"
	var p *string = &creature

	fmt.Println(creature)
	fmt.Println(p)
	fmt.Println(*p)
	fmt.Println("\n")

	*p = "jellyfish"
	fmt.Println(*p)
	fmt.Println(creature)
}
