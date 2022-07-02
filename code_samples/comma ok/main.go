package main

import "fmt"

func main() {

	c := make(chan int)
	go func() {
		c <- 42
		close(c)
	}()
	v, ok := <-c
	fmt.Println(v, ok)
 // gives an output, 0 for c channel. 
	v, ok = <-c
	fmt.Println(v, ok)

}
