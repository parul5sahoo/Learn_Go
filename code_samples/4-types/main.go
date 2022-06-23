package main

import ( "fmt"
         "runtime"
)

var a int
// creating our own types
// hotdog is a custom type with int as underlying type
type hotdog int
var b hotdog

 var x int
 var y float64
 var z int8 = 127

func main() {
	a = 42
	b = 43
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	// Conversion, which is called casting in other langs
	a = int(b)
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	x = 43
	y = 42.42345
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", z)



}