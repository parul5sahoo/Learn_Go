package main

import "fmt"

var x int
var y string
var z bool

func main() {
	x = 42
	y = "Jame Bond"
	z = true

	//using short declar
	// x := 42
	// y := "Jame Bond"
	// z := true

	//Printing them using single print stmt
	fmt.Println(x, y, z)
	//Printing them using multiple print stmtss
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

}