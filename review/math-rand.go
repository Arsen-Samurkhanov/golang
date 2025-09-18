package main

import "math/rand"

func main() {
	for i := 0; i < 10; i++ {
		println(rand.Intn(100))
	}
}
