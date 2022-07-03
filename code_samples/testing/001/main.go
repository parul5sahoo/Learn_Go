package main

import (
	"fmt"

)

func mySum(xi ...int) int {
	sum := 0 
	for _, v := range xi {
		sum += v
	}
	return sum
	// for the error and test to fail
	// return sum+1
} 

func main() {

	fmt.Println("2 + 3 =", mySum(2, 3))
	fmt.Println("7 + 4 =", mySum(7, 4))
	fmt.Println("9 + 5 =", mySum(9, 5))
}