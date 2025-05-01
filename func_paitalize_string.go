package main

import (
	"fmt"
	"strings"
)

func capitalize(name string) string {
	return strings.ToTitle(name)
}

func main() {
	name := "arsen"
	fmt.Println(capitalize(name))

}
