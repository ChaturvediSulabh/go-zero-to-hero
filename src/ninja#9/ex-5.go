// Fix the race condition you created in exercise #4 by using package atomic
package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("CPUs: ", runtime.NumCPU())
	fmt.Println("GOROUTINEs: ", runtime.NumGoroutine())
	var incr int64
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			atomic.AddInt64(&incr, 1)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(incr)
}
