package main

import "fmt"

type person struct {
	name string
}

func main() {
	p1 := person{
		name: "Parul Sahoo",
	}
	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)
}

func changeMe(p *person) {
	p.name = "Miss Parul"
	(*p).name = "Miss ParulS"
	// diff ways to derefence the address the struct
}
