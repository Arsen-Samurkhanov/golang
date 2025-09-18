package main

import "fmt"

type Creature struct {
	Name     string
	Greeting string
}

func (c Creature) Greet() {
	fmt.Printf("%s says %s\n", c.Name, c.Greeting)

}

func main() {
	sammy := Creature{
		Name:     "Sammy",
		Greeting: "Hello!",
	}

	etien := Creature{
		Name:     "Etien",
		Greeting: "Salute!",
	}
	sammy.Greet()
	etien.Greet()
	//Creature.Greet(sammy)

}
