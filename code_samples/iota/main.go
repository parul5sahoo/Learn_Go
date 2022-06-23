package main

import (
	"fmt"
)

// const a = 42
// const b = 42.78
// const c = "Parul Sahoo"

// or one can declare constants as such
const (
	a = iota + 1
	b = iota
	c = iota
)

const (
	d = iota + 1
	e
	f
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Printf("%T\n", d)
	fmt.Printf("%T\n", e)
	fmt.Printf("%T\n", f)

}
