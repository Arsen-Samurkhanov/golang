package main

import "fmt"

type Creature struct {
	Species string
}

func (c *Creature) Reset() {
	c.Species = ""

}

func main() {
	var creature Creature = Creature{Species: "shark"}
	fmt.Printf("%+v\n", creature)
	creature.Reset()
	fmt.Printf("%+v\n", creature)

}
