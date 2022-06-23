package main

import "fmt"

func main() {
	x := [10]int{32, 43, 63, 48, 49, 10, 52, 38, 91, 100}

	for i, v := range x {
		fmt.Printf("value at index %v is %d of type %T\n", i, v, v)
	}

}
