package main

import "fmt"

func sum_mult(num int) (int, int) {
	return num + num, num * num

}

func main() {
	num := 10
	sum, mult := sum_mult(num)

	fmt.Println(sum)
	fmt.Println(mult)

}
