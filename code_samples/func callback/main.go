package main

import "fmt"

func main() {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := sum(ii...)
	fmt.Println("sum of all numbers", s)

	s2 := even(sum, ii...)
	fmt.Println("sum of all even numbers", s2)

	s3 := odd(sum, ii...)
	fmt.Println("sum of all odd numbers", s3)

}

func sum(xi ...int) int {
	fmt.Printf("%T \n", xi)
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

// if we just want to print the sum of even nums

func even(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}

func odd(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 != 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}
