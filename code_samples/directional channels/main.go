package main

import "fmt"

// func main() {
// 	c := make(chan <- int, 2) // a send only channel i.e only vals can be send to the channel and cannot be received.

// 	c <- 42
// 	c <- 43

// 	fmt.Println(<-c) // throws err  invalid operation: <-c (receive from send-only type chan<- int)
// 	fmt.Println(<-c)
// 	fmt.Println("--------")
// 	fmt.Printf("%T\n", c)
// }

// learning Receive only channel
func main() {
	c := make(<- chan int, 2) // a receive only channel i.e only be received from the channel and cannot vals can be send to the channel.

	// c <- 42 // throws an err invalid operation: c <- 42 (send to receive-only type <-chan int
	// c <- 43

	fmt.Println(<-c) 
	fmt.Println(<-c)
	fmt.Println("--------")
	fmt.Printf("%T\n", c)


}