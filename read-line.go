package main

import (
	"fmt"
)

func main() {
	fmt.Println("Please enter your name")
	var name string
	fmt.Scanln(&name)
	fmt.Println("Hi, %s! Im Go", name)

}
