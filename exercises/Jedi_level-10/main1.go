package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	
	// using func literal
	go func(c chan<- int) {
		c <- 42
	}(c)

	func(c <-chan int) {
		fmt.Println("value of chan is:", <-c)
	}(c)

	fmt.Println("about to exit")

}


// package main

// import (
// 	"fmt"
// )

// func main() {
// 	// using buffered channels
// 	c := make(chan int, 1)
// 	c <- 42
// 	fmt.Println("value of chan is:", <-c)
// 	fmt.Println("about to exit")

// }




