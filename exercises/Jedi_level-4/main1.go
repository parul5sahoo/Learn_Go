package main

import "fmt"

func main() {
	// initializing an array with composite literal
	x := [5]int{32, 43, 63, 48, 49}

	for i, v := range x {
		fmt.Printf("value at index %v is %d of type %T\n", i, v, v)
	}

}
