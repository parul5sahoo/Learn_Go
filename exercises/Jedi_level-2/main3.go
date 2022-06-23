package main

import "fmt"

const(
	// untyped const
	a = 42
	//typed const
	b int = 43
)

func main() {
	fmt.Println(a, b)
}