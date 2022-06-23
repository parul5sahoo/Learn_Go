package main

import "fmt"

func main() {
	for i := 10; i <= 100; i++ {
		x := i / 4
		y := i % 4
		fmt.Printf("The remainder of %v divide by 4 is %v and quotient is %v\n", i, x, y)
		// or to lessen the no. of vars used 
		// fmt.Printf("The remainder of %v divide by 4 is %v and quotient is %v\n", i, i%4, i/4)

	}
}
