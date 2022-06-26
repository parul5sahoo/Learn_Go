package main

import "fmt"

func main() {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	n := foo(ii...)
	// we need unfurl the slice

	ii2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	m := bar(ii2) // we do not need to unfurl it
	fmt.Printf("The sum of slice 1 is %v", n)
	fmt.Printf("The sum of slice 2 is %v", m)
}

func foo(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func bar(xi []int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}
