package main

import "fmt"

func main() {
	sammy := map[string]string{"name": "Sammy", "animal": "shark", "color": "blue"}
	fmt.Println(sammy)
	fmt.Println("")
	fmt.Println(sammy["color"])
}
