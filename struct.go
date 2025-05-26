package main

import "fmt"

type Octopus struct {
	Name  string
	Color string
}

func main() {

	oct := Octopus{
		Name:  "Jessy",
		Color: "Orange",
	}
	fmt.Println(oct)
	fmt.Println(oct.Name)
	fmt.Println(oct.Color)

}
