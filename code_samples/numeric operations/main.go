package main

import "fmt"

func main() {

	a := 7.0 / 3
	b := 7.0 / 3
	c := 7 / 3
	d := 7 / 3.0
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(a == b)
	fmt.Println(b == d)
	// these throw err
	// invalid operation: a == c (mismatched types float64 and int)
	// fmt.Println(a == c)
	// fmt.Println(c == d)
}
