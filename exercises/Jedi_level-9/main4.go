package main

import (
	"fmt"
	"runtime"
)

func main(){

	fmt.Println("OS\t", runtime.GOOS)
	fmt.Println("Arch\t", runtime.GOARCH)
	fmt.Println("Go root\t", runtime.GOROOT())
	

}