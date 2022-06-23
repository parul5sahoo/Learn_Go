package main

import "fmt"

type new_type int

var x new_type


func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	
	
	
}