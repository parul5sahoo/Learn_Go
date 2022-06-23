package main

import (
	"fmt"
)

func main() {
	x := 84
	y := x << 1
	fmt.Printf("%T\t%d\t%b\t%#X\n", x, x, x, x)
	fmt.Printf("%T\t%d\t%b\t%#X\n", y, y, y, y)

}
