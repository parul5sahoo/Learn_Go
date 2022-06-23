package main

import "fmt"

func main() {
	// slices are declared using composite literal i.e
	// var := type{values} for e.g.
	x := []int{1, 2, 3, 4, 57}
	fmt.Println(x)
	fmt.Println(len(x))
	// slices help us group values of the same type.
	for i, v := range x {
		fmt.Printf("value at index %v is %v\n", i, v)
	}
	// slicing a scile with : operator
	fmt.Println(x[1:])
	// prints the slice starting from  index 1
	fmt.Println(x[1:3])
	// prints the slice starting from  index 1 and ending at index 2 excluding 3.

	for i := 0; i < len(x); i++ {
		fmt.Println(x[i])
	}
	// appending values to a slice
	x = append(x, 77, 88, 99, 1014)
	fmt.Println(x)
	// appending slice to a slice using variading parameter ...
	y := []int{23, 46, 91}
	x = append(x, y...)
	fmt.Println(x)
	// deleteing values from a slice using variadic params
	x = append(x[:2], x[4:]...)
	fmt.Println(x)

	z := make([]int, 10, 12)
	fmt.Println(z)
	fmt.Println(len(z))
	fmt.Println(cap(z))
	z[0] = 42
	z[9] = 999
	fmt.Println(z)
	fmt.Println(len(z))
	fmt.Println(cap(z))
	z = append(z, 3432)
	fmt.Println(z)
	fmt.Println(len(z))
	fmt.Println(cap(z))
	z = append(z, 3434)
	fmt.Println(z)
	fmt.Println(len(z))
	fmt.Println(cap(z))
	// since the capacity is exceeded a new array is created and the old arrays values are copied to the new one and the capacity doubles.
	z = append(z, 3435)
	fmt.Println(z)
	fmt.Println(len(z))
	fmt.Println(cap(z))
	
	// multi-dimensional slices
	jb := []string{"James", "Bond", "Chocolate", "Martini"}
	fmt.Println(jb)
	mp := []string{"Miss", "MoneyPenny", "Strawberry", "hazelnut"}
	fmt.Println(mp)

	xp := [][]string{jb, mp}
	fmt.Println(xp)

}
