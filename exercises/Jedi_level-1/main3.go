package main

import "fmt"

var x = 42
var y = "James Bond"
var z = true

func main() {
	s := fmt.Sprintf("%d\t%s\t%t\n", x, y, z)
	// another way is 
	// s := fmt.Sprintf("%v\t%v\t%v\n", x, y, z)
	fmt.Println(s)
	
}