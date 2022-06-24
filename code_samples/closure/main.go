package main

import "fmt"

func main() {
	a := incrementor()
	b := incrementor()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())

	{
		y := 42
		fmt.Println(y)
	}
	// will throw an err coz scope of y is until that code block only.
	// fmt.Println(y)
}

func incrementor() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}
