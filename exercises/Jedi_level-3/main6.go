package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("not true shouldn't print")
	case true:
		fmt.Println("is true should print")
	}

}
