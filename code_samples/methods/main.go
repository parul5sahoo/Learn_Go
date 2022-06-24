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

// besically a way of saying that any value of type secretAgent has access to func speak
// and the method is accesses via ./dot notation.
func (s secretAgent) speak() {
	fmt.Println("I am ", s.first, s.last)
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
		},
		ltk: true,
	}

	fmt.Println(sa1)
	sa1.speak()
}
