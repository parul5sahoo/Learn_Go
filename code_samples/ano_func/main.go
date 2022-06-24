package main

import "fmt"

func main() {
	var x int
	x = 42
	foo()
	// anonymous functions
	func() {
		fmt.Println("Anonymous func ran")
	}()

	func(x int) {
		fmt.Println("The meaning of life:", x)
	}(x)// these params at the end pass the params. so an empty param stands for empty interfaces.

}

func foo() {
	fmt.Println("foo ran")
}
