package main

import (
	"fmt"
)

const (
	_  = iota // to have the 0 value of iota ignored by the compiler
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
	// within the parens the iota increments and
	//we cannot use the short declar format outside the code block
)

func main() {
	// basic bit shifting
	// x := 4
	// fmt.Printf("%d\t\t%b\n", x, x)
	// y := x << 1
	// fmt.Printf("%d\t\t%b\n", y, y)

	// to see iota in action shifting of 10 places

	// kb := 1024
	// mb := 1024 * kb
	// gb := 1024 * mb
	fmt.Printf("%d\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t%b\n", mb, mb)
	fmt.Printf("%d\t%b\n", gb, gb)

}
