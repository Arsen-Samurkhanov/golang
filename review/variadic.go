package main

import "fmt"

func main() {
	sayHello("Sammy")
	sayHello("SAmmy", "Jessica", "Drew", "Jamie")
}

func sayHello(names ...string) {
	for _, n := range names {
		fmt.Printf("Hello %s\n", n)
	}

}
