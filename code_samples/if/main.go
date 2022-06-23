package main

import "fmt"

func main() {
	// using pre-fefined constyants
	if true {
		fmt.Println("001")
	}
	if false {
		fmt.Println("002")
	}
	if !true {
		fmt.Println("003")
	}
	if !false {
		fmt.Println("004")
	}
	// using conditional statements
	if 2 == 2 {
		fmt.Println("005")
	}
	if !(2 == 2) {
		fmt.Println("006")
	}
	if 2 != 2 {
		fmt.Println("007")
	}
	if !(2 != 2) {
		fmt.Println("008")
	}
	// using initialisation statement
	if x := 42; x == 2 {
		fmt.Println("Yes")
	}
	fmt.Println("here's a stmnt")
	fmt.Println("somethn else")
	// fmt.Println(x) will not run since the scope of x 
	// is limited to the if block only
}
