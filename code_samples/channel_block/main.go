package main

import "fmt"

// func main() {
// 	c := make(chan int)
// 	// this wont run as it is and throw err:- fatal error: all goroutines are asleep - deadlock!
// 	//goroutine 1 [chan send]:

// 	// so we add this to create a new go routine to run concurrently
// 	go func(){
// 		c <- 42
// 	}()
	
// 	fmt.Println(<-c)

// }

// A second method is 

// func main() {
// 	c := make(chan int, 1) // this is a buffer channel, so it allows one val to sit in there.
	
// 	
// 		c <- 42
// 	
	
// 	fmt.Println(<-c)

// }


// Writing an unsuccesful buffer

// func main() {
// 	c := make(chan int, 1) // this is a buffer channel, so it allows one val to sit in there.
	
// 		c <- 42
// 		c <- 43 // the buffer has limit for one and hence the code throws an err.
	
// 	fmt.Println(<-c)

// }

// making the above code run successfully

func main() {
	c := make(chan int, 2) // this is a buffer channel, so it allows two vals to sit in there.
	
		c <- 42
		c <- 43 
	
	fmt.Println(<-c)
	fmt.Println(<-c)

}