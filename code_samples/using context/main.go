package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	fmt.Println("erros check 1:", ctx.Err())
	fmt.Println("num gortins 1:", runtime.NumGoroutine())

	go func() {
		n := 0
		for {
			select {
			case <-ctx.Done():
				return
			default:
				n++
				time.Sleep(time.Millisecond * 200)
				fmt.Println("working", n)
			}
		}
	}()
	time.Sleep(time.Second * 2)
	fmt.Println("error check 2:", ctx.Err())
	fmt.Println("num of gortins 2:", runtime.NumGoroutine())

	fmt.Println("about to cancel context")
	cancel()
	fmt.Println("cancelled context")

	time.Sleep(time.Second * 2)
	fmt.Println("error check 3:", ctx.Err())
	fmt.Println("num of gortin 3:", runtime.NumGoroutine())
}

// Another example  

// package main

// import (
// 	"context"
// 	"fmt"
// )

// func main() {

// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel() // cancel when we are finished

// 	for n := range gen(ctx) {
// 		fmt.Println(n)
// 		if n == 5 {
// 			break
// 		}
// 	}
// }

// func gen(ctx context.Context) <-chan int {
// 	dst := make(chan int)
// 	n := 1

// 	go func() {
// 		for {
// 			select {
// 			case <-ctx.Done():
// 				return // returning not to leak the goroutine
// 			case dst <- n:
// 				n++
// 			}
// 		}
// 	}()
// 	return dst
// }
