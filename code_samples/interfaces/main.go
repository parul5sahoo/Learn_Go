package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

func (s secretAgent) speak() {
	fmt.Println("I am ", s.first, s.last, "I was called in secretAgent")
}

func (p person) speak() {
	fmt.Println("I am ", p.first, p.last, "I was called in human")
}

type human interface {
	// any type that has a method speak is of the type human as well
	speak()
}

type hotdog int

func bar(h human) {
	switch h.(type) {
	// This is an example assertion
	case person:
		fmt.Println("I was passed into bar as person", h.(person).first)
	case secretAgent:
		fmt.Println("I was passed into bar as secretAgent", h.(secretAgent).first)
	}
	fmt.Println("I am a human too")
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
		},
		ltk: true,
	}

	sa2 := secretAgent{
		person: person{
			first: "Harold",
			last:  "Jane",
		},
		ltk: true,
	}

	p1 := person{
		first: "Dr.",
		last:  "Maisy",
	}

	fmt.Println(sa1)
	sa1.speak()
	sa2.speak()
	p1.speak()
	bar(sa1)
	bar(sa2)
	bar(p1)

	// Example conversion
	var x hotdog = 42
	fmt.Printf("The value of x is %v it is of type %T\n", x, x)
	var y int = int(x)
	fmt.Printf("The value of y is %v it is of type %T\n", y, y)
}
