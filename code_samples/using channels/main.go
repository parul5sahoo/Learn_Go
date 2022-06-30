package main

import "fmt"

func main() {

	c := make(chan int)
	// send
	go foo(c)

	// receive
	bar(c)
	// removing go before bar will make the bar func wait until the value is sent to the channel. 
	// otherwise it won't wait for the value to be sent and execute it already.

	fmt.Println("about to exit")
}

// send
func foo(c chan<- int){
	c <- 42
}

// receive
func bar(c <-chan int){
	fmt.Println("value of chan", <-c)
}