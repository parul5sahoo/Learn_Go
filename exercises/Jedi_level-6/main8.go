package main

import "fmt"

func main() {
	f := golang()
	fmt.Printf("And it returned a value %v", f())
}

func golang() func() int {
	return func() int {
		fmt.Println("This is the returned func")
		return 42
	}

}
