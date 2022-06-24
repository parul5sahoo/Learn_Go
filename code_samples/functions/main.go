package main

import "fmt"

func main() {
	defer foo()
	bar("Parul")
	s1 := woo("Nidhi")
	fmt.Println(s1)
	x, y := mouse("Nirmal", "Hassan")
	fmt.Println(x)
	fmt.Println(y)
	Sum := sum(2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("The sum total is ", Sum)

	/*
		xi := []int{2, 3, 4, 5, 6, 7, 8, 9} // will work
	    Sum := sum(xi...) // this is unfurling basically passing the same slice that the func will be using to store int values.
	    however doing Sum := sum(xi) would throw an err, []int if diift. from ...int
		x := sum() would also work because variadic param means 0 or more values of a specific type.
		
	*/
}
func foo() {
	fmt.Println("Hello from foo")
}

func bar(s string) {
	fmt.Println("Hello,", s)
}

func woo(s string) string {
	return fmt.Sprint("Hello from woo, ", s)
}
func mouse(fn string, ln string) (string, bool) {
	a := fmt.Sprint(fn, " ", ln, `, says "Hello"`)
	// default val of bool is false
	var b bool
	return a, b
}

func sum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for i, v := range x {
		fmt.Println("for the item in indes position ", i, " we are now adding, ", v, " to the total which is now ", sum)
		sum += v
	}
	return sum
}