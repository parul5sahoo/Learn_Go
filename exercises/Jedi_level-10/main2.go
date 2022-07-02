package main

import (
	"fmt"
)

func main() {
	cs := make(chan int)

	go func() {
		cs <- 42
	}()
	fmt.Println(<-cs)
	fmt.Printf("\t cs type is %T\n", cs)

}

// example :-2

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	cr := make(chan int)

// 	go func() {
// 		cr <- 42
// 	}()
// 	fmt.Println(<-cr)
// 	fmt.Println("--------\n")
// 	fmt.Printf("cr type is %T\n", cr)

// }

