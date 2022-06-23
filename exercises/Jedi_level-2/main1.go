package main

import "fmt"



func main() {
	a := 42
	// always keep your scopes as narrow as possible.
	fmt.Printf("%d\t%b\t%#x\n", a, a, a)
	

}