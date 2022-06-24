package main

import "fmt"

func main() {
	n := factorial(5)
	fmt.Println(n)
	nloop := loopFact(5)
	fmt.Println(nloop)
}

func factorial(i int) int {
	if i == 0 {
		return 1
	}
	return i * factorial(i-1)
}

func loopFact(n int) int {
	total := 1
	i := 1
	// for loop w/o an initializer w/o needing to create a new variable
	for ; n > 0; n-- {
		total *= i
		i++
	}
	return total
}
