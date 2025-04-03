package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := strconv.Itoa(12)
	fmt.Printf("%q\n", a+a)
	fmt.Println(fmt.Sprint(45.15))
}
