package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Printf("Hello I am %v %v and my age is %v", p.first, p.last, p.age)
}

func main() {

	p1 := person{
		first: "Parul",
		last:  "Sahoo",
		age:   21,
	}

	p1.speak()

}
