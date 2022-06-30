package main 

import "fmt"

type person struct {
	first string

}

type human interface {
	speak()
}

func (p *person) speak() {

	fmt.Println("Hello")
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person {
		first: "James",
	}
	// saySomething(p1)
	// won't work coz speak method receives pointer to a person hence
	saySomething(&p1) // this works
	// and this works because of speak method 
	p1.speak() // coz the func can be called by pointer to a person or person
}