package main

import (
	"fmt"
)

func main() {
	s := "Hello, World"
	// using raw string literal with backticks
	t := `"Hello, 

World with spaces"`
	fmt.Println(s)
	fmt.Printf("%T\n", s)
	fmt.Println(t)
	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U\t", s[i])
	}
	
	fmt.Println("")

	for i, v := range s {
		fmt.Println(i, v)
	}

	// to view hexadecimals

	for i, v := range s {
		fmt.Printf("at position %d we have hex %#x", i, v)
		fmt.Println("")
	}

}
