package main

import (
	"fmt"
	"math"
)

func main() {
	s := 80
	t := 6
	fmt.Println(s / t)
	fmt.Println(float32(s) / float32(t))
	fmt.Println(s % t)

	q := 36.0
	r := 8.0
	fmt.Println(math.Mod(q, r))
}
