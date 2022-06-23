package main

import (
	"fmt"
)

// const a = 42
// const b = 42.78
// const c = "Parul Sahoo"

// or one can declare constants as such
const (
	a = 42
	b = 42.78
	c = "parul Sahoo"
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)

}
