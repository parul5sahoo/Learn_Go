package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// recreating race condns & fixing it using mutex and using atomic.
// var mu sync.Mutex

func main() {
	var wg sync.WaitGroup

	var incrementor int64
	gs := 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			// mu.Lock()
			atomic.AddInt64(&incrementor, 1)
			// v := incrementor
			// one can remove gosched while using mutex
			// runtime.Gosched()
			// v++
			// incrementor = v
			
			// fmt.Println(incrementor)

			 fmt.Println(atomic.LoadInt64(&incrementor))
			// mu.Unlock()
			wg.Done()
		}()

	}
	wg.Wait()
	fmt.Println("end value: ", incrementor)
}
