package main

import "fmt"

const (
	year     = 365        //untyped constant
	leapYear = int32(366) //typed constant
)

func main() {
	hours := 24
	minutes := int32(60)

	fmt.Println(hours * year)
	fmt.Println(minutes * year)
	fmt.Println(minutes * leapYear)
}
