package main

import (
	"fmt"
	"strings"
)

func main() {
	hello()
	names()

}

func hello() {
	fmt.Println("Hello World!")
}

func names() {
	fmt.Println("Enter your name:")
	var name string
	fmt.Scanln(&name)
	for _, v := range strings.ToLower(name) {
		if v == 'a' || v == 'e' || v == 'i' || v == 'o' || v == 'u' {
			fmt.Println("Your name contains a vowel.")
			return
		}
	}
	fmt.Println("Your name does not contains a vowel.")

}
