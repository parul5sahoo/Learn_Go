package main

import "fmt"

func main() {
	f := foo(3, 5)
	sum, strn := bar(1, 3)
	fmt.Printf("The sum of int is %v and %v %v ", f, strn, sum)
}

func foo(i int, t int) int {
	s := i + t
	return s
}

func bar(i int, t int) (int, string) {
	str := fmt.Sprintf("The 2 values passed to bar are %v and %v and their sum is ", i, t)
	s := i + t
	return s, str
}
