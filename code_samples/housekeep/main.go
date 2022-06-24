package main

import "fmt"

type foo int

var x foo

const bar foo = 42

func main() {
	x = 44
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", int(x))
	fmt.Println(x)
	fmt.Printf("%T\n", bar)
	fmt.Println(bar)

	var y int
	// y = bar
	// will return an error because bar is not const ot ype kind
	// but of type int foo anf var y is of type int and even though for
	// foo the underlying type is int parallel types cannot be assigned ot consts.
	fmt.Println(y)
}
