package main

import "fmt"

func main() {

	coral := [4]string{"blue coral", "staghorn coral", "pillar coral", "elkhorn coral"}

	fmt.Println(coral)

	fmt.Printf("%q\n", coral)

	fmt.Println(coral[1])

	coral[1] = "foloise colar"

	fmt.Println(coral[1])

	fmt.Println(len(coral))

}
