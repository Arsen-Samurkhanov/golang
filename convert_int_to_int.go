package main

import "fmt"

func main() {
	var index int8 = 15
	var bigIndex int32

	bigIndex = int32(index)

	fmt.Println(bigIndex)
	fmt.Printf("index data type %T\n", index)
	fmt.Printf("bigIndex data type %T\n", bigIndex)

}
