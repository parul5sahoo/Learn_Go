package main

import "fmt"

func main() {
	a := 42
	fmt.Println(a)
	fmt.Println(&a)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)

	fmt.Println(" ")

	var b *int = &a
	fmt.Println(b)
	fmt.Println(*b)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", *b)

	fmt.Println(" ")

	c := &b
	fmt.Println(c)
	fmt.Println(*c)
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", *c)

	fmt.Println(" ")

	*b = 43 // changing the value stored at a ceratin address using * operator.
	fmt.Println(a)


	x := 0
	fmt.Println(x)
	foo(&x)
	fmt.Println(x)

}

func foo(y *int) {
	fmt.Println(y)
	fmt.Println(*y)
	*y = 43
	fmt.Println(y)
	fmt.Println(*y)
}