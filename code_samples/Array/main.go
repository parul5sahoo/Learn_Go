package main

import "fmt"

func main() {
	var x [5]int
	var y [6]int // is a totally diff. type of array than x because of diff length.
	fmt.Println(x)
	x[3] = 42
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(len(x))

}
