package main

import (
	"fmt"
	"sort"
)

func main() {
	// sorting slices
	xi := []int{4, 7, 3, 42, 99, 18, 16, 56, 12}
	xs := []string{"James", "P", "S", "Narul Prusty", "Dr.Strange"}

	fmt.Println(xi)
	fmt.Println(xs)

	sort.Ints(xi)
	sort.Strings(xs)
	fmt.Println("--------", "now sorted")

	fmt.Println(xi)
	fmt.Println(xs)

}
