package main

import "fmt"

func main() {

	sammy := map[string]string{"name": "Sammy", "animal": "shark", "color": "blue"}
	fmt.Println(sammy)
	fmt.Println(sammy["name"])
	fmt.Println(sammy["animal"])
	fmt.Println(sammy["color"])

	fmt.Println("\n")

	//   iterate map
	for key, value := range sammy {
		fmt.Printf("%q is the key for value %q\n", key, value)
	}

	fmt.Println("\n")

	//    print only keys
	keys := []string{}
	for key := range sammy {
		keys = append(keys, key)
	}
	fmt.Printf("%q", keys)

	fmt.Println("\n")

	//print only values
	items := make([]string, len(sammy))
	var i int
	for _, v := range sammy {
		items[i] = v
		i++
	}
	fmt.Printf("%q", items)

	fmt.Println("\n")

	// number of items
	fmt.Println(len(sammy))

}
