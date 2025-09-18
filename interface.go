package main

import "fmt"

type Article struct {
	Title  string
	Author string
}

func (a Article) String() string {
	return fmt.Sprintf("The %q article was written by %s.", a.Title, a.Author)
}

type Book struct {
	Title  string
	Author string
	Pages  int
}

func (b Book) String() string {
	return fmt.Sprintf("The %q book was written by %s.", b.Title, b.Author)
}

type Stringer interface {
	String() string
}

func main() {
	a := Article{
		Title:  "Ubderstanding I nterfaces in Go",
		Author: "Sammy Shark",
	}
	Print(a)
	//fmt.Println(a.String())

	b := Book{
		Title:  "All about Go",
		Author: "Johny Dolphin",
		Pages:  25,
	}
	Print(b)
	//fmt.Println(b.String())
}

func Print(s Stringer) {
	fmt.Println(s.String())
}
