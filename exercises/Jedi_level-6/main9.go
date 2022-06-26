// package main

// import "fmt"

// func main() {
// 	s := 93
// 	var res = bar(foo, s)
// 	fmt.Printf("The result of bar(s) is %v", res)

// }

// func foo(i int) int {
// 	return 91 + i
// }

// func bar(f func(i int) int, i int) int {
// 	return f(i)
// }


// or another example of passing a func in a func is 

package main

import "fmt"

func main() {

	g := func(xi []int) int {
		if len(xi) == 0 {
			return 0
		}
		if len(xi) == 1 {
			return xi[0]
		}
		return xi[0] + xi[len(xi)-1]
	}

	x := foo(g, []int{1, 2, 3, 4, 5})
	fmt.Println(x)
}

func foo(f func(xi []int) int, ii []int) int {
	n := f(ii)
	n++
	return n
}
