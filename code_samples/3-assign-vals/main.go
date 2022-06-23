package main

import "fmt"

var y = 42

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Printf("%b\n", y)
	fmt.Printf("%x\n", y)
	fmt.Printf("%#x\n", y)
	// or you can use the formatting in one line
	fmt.Printf("%#X\t%b\t%x\t\n", y, y, y)
	// or you can add it to a string too using Sprint
	// Here \n \t and stuff are escape chars
	s := fmt.Sprintf("%#X\t%b\t%x\t\n", y, y, y)
    fmt.Println(s)
	fmt.Printf("%v", y)
}