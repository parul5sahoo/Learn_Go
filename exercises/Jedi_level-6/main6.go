package main

import (
	"fmt"
)

func main() {
	x := 42
	func(x int) {
		fmt.Printf("hey my age is %v", x)
	}(x)
 // anonymous self-executing code.
    
}

//