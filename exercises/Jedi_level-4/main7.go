package main

import "fmt"

func main() {
	xs1 := []string{"James", "Bond", "Shaken, not stirred"}
	xs2 := []string{"Miss", "MoneyPenny", "Hello, James."}
	fmt.Println(xs1)
	fmt.Println(xs2)
	xs3 := [][]string{xs1, xs2}
	fmt.Println(xs3)

	for i, xs := range xs3 {
		fmt.Println("record:", i)
		for j, val := range xs {
			fmt.Printf("\t indes position: %v \t value: \t %v\n", j, val)
		}
	}

}
