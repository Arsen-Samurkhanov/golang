package main

import "fmt"

func main() {
	fmt.Println(double(8))

}

func double(x int) int {
	return x * 2
}
