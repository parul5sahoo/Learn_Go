package main

import (
	"fmt"
	// "dog"
)

type canine struct {
	name string
	age uint
}

func main() {
	fido := canine {
		name: "Fido",
		// age: dog.Years(10),
	}
	fmt.Println(fido)
	// we need to create a go doc and use it
}

