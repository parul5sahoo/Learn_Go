package main

import "fmt"

func main() {
	// WE Can assign a function to a variable.
	f := func() {
		fmt.Println("my first func expression")
	}
	f()

	g := func(x int) {
		fmt.Println("The yr I was born in:", x)
	}
	g(2002)
}
