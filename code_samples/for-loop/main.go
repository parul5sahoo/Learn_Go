package main

import "fmt"

func main() {
	// syntax for "for" with Forclause, for init; condition; post{}
	for i := 0; i <= 100; i++ {
		fmt.Println(i)
	}
	// nested loops
	for i := 0; i <= 100; i++ {
		fmt.Printf("The outer loop: %d\n", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t\tThe inner loop: %d\n", j)
		}
	}
	// for loop with single condition basically while stmt
	x := 1

	for x < 10 {
		fmt.Println(x)
		x++
	}

	// if within for and usage of break & continue
	for {
		x++
		if x > 100 {
			break
		}
		if x%2 != 0 {
			continue
		}
		fmt.Println(x)

	}

	// for loop with range clause



}
