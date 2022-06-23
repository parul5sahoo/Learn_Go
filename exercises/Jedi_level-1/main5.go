package main

import "fmt"

type new_type int

var x new_type
var y int


func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
		
}