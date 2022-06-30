package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {

	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	var	counter int64

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

		// Mutex is used to lock down a code block to prevent race conditions, which arise when multiple go routines are trying access the same shared vars.
		// There's also the flexibility to lock a code from reading or writing using the RWMutex, RLocker() or RLock() or RUnlock()
	

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			// gosched can be avided
			runtime.Gosched()
			fmt.Println("Counter\t", atomic.LoadInt64(&counter))
			// time.Sleep(time.Second * 2)
			
			
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("count:", counter)
}
