package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("begin gs", runtime.NumGoroutine())
	fmt.Println("begin CPU", runtime.NumCPU())
	fmt.Println("Hello from main")
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		fmt.Println("Hello from func 1")
		wg.Done()
	}()
	go func() {
		fmt.Println("Hello from func 2")
		wg.Done()
	}()
	fmt.Println("mid gs", runtime.NumGoroutine())
	fmt.Println("mid CPU", runtime.NumCPU())
	wg.Wait()

	fmt.Println("about to exit")
	fmt.Println("end gs", runtime.NumGoroutine())
	fmt.Println("end CPU", runtime.NumCPU())
}
