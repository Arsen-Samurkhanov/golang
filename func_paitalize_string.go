package main

import (
	"fmt"
	"strings"
)

func capitalize(name string) string {
	return strings.ToTitle(name)
}

func main() {
	var name string
	fmt.Println("Emter your name in lower case")
	fmt.Scanln(&name)
	fmt.Println(capitalize(name))

}
