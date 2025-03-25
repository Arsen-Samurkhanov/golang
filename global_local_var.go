package main

import "fmt"

var g = "global"

func prinLocal() {
	l := "local"
	fmt.Println(l)
}

func main() {
	prinLocal()
	fmt.Println(g)
}
