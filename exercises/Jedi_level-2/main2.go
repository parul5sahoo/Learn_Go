package main

import (
	"fmt"
)

func main() {
	x := 1
	y := 20
	z := x == y
	a := x >= y
	b := x <= y
	c := x << 2
	d := y >> 1
	e := x > y
	f := x < y
	fmt.Println(x, y, z, a, b, c, d, e, f)

}
