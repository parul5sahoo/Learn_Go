package main

import (
	"fmt"
)

var g func()

func main() {

	f := func() {
		for i := 0; i < 3; i++ {
			fmt.Println(i)
		}

	}
	// anonymous self-executing code.
	g = f
	g()
	fmt.Printf("This exec is done by g %T\n", g)
	fmt.Println("done")
}
