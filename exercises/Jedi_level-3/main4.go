package main

import "fmt"

func main() {
	x := 2002
	for {
		if x <= 2022 {
			fmt.Println(x)
		} else {
			break
		}
		x++
	}
}
