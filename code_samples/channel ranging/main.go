package main

import "fmt"

// func main() {

// 	c := make(chan int)
// 	// send
// 	go foo(c)

// 	// receive
// 	for v := range c {
// 		fmt.Println(v)
// 	}

// 	fmt.Println("about to exit")
// }

// // send
// func foo(c chan<- int){
// 	for i := 0; i< 100; i++{
// 		c <- i
// 	}
// 	close(c)
// }

// cleaner code represenataion

func main() {

	c := make(chan int)
	// send
	go func(){
		for i := 0; i< 100; i++ {
			c <- i
		}
		// we need to close the channel otherwise it keeps waiting for the value event after the loop ends
		// and that throws and err all goroutines are asleep - deadlock!

		close(c)
	}()

	// receive
	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}