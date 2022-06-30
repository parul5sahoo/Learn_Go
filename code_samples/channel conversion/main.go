package main

import "fmt"

func main() {
	c := make(chan int)
	cs := make(chan<- int) // send only
	cr := make(<-chan int) // receive only

	fmt.Println("-------")
	fmt.Printf("c\t%T\n", c)
	fmt.Printf("cr\t%T\n", cr)
	fmt.Printf("cs\t%T\n", cs)

	// specific to general doesn't assign
	// c = cr
	// c = cs

	// checking if the reverse of true
	cr = c
	fmt.Println("-------")
	fmt.Printf("cr\t%T\n", cr)
	cs = c
	fmt.Println("-------")
	fmt.Printf("cs\t%T\n", cs)

	// changing the type of generic channel to specific channel
	// using type conversion

	fmt.Println("-------")
	fmt.Printf("c\t%T\n", (<-chan int)(c))
	fmt.Println("-------")
	fmt.Printf("c\t%T\n", (chan<- int)(c))

	// changing the type of specific channel to one another
	// throws err
	// fmt.Println("-------")
	// fmt.Printf("cs\t%T\n", (<-chan int)(cs)) // cannot convert cs (variable of type chan<- int) to type <-chan int
	// fmt.Println("-------")
	// fmt.Printf("cr\t%T\n", (chan<- int)(cr))

	// changing the type of specific channel to generic
	// throws err
	// fmt.Println("-------")
	// fmt.Printf("cs\t%T\n", (chan int)(cs)) // cannot convert cs (variable of type chan<- int) to type chan int
	// fmt.Println("-------")
	// fmt.Printf("cr\t%T\n", (chan int)(cr))

}
