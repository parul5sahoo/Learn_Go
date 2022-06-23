package main

import (
	"fmt"
)

const (
	// pritning the last 4 yrs using iota..
	x = 2018 + iota
	y = x + iota
	z = x + iota
	a = x + iota
)

func main() {

	fmt.Println(x, y, z, a)

}
