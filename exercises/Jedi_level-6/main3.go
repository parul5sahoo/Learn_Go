package main

import "fmt"

func main() {
	defer foo()
	fmt.Println("I am good at programming")
	fmt.Println("I can definitely crack GSoC next year")
}

func foo() {

	defer func() {
		fmt.Println("coz I am determined and unstoppable")
	}()
	str := "I will work hard enough for it.."
	fmt.Println(str)
}

func bar(xi []int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}
