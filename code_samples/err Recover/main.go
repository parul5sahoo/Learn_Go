package main 

import "fmt"

func main() {
	f()
	fmt.Println("Returned normally from f.")
}

func f() {

	defer func() {
		// a common idiom to create an idiom to assign a val to a 
		// closed scope var and add an if condition...
		if r := recover(); r != nil {
			fmt.Println("Recovered from f", r)
		}
	}()

	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}