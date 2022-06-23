package main

import "fmt"

func main() {
	// to print the unicodes from 33 - 122

	for i := 33; i <= 122; i++ {
		fmt.Printf("%v\t%#x\t%#U\n", i, i, i)

	}
}
