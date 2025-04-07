package main

import "fmt"

var x int64 = 57
var y float64 = float64(x)

func main() {
	fmt.Printf("%.3f\n", y)
}

