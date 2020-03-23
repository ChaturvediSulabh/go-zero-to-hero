// Fix the race condition you created in the previous exercise by using a mutex
// it makes sense to remove runtime.Gosched()
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs: ", runtime.NumCPU())
	fmt.Println("GOROUTINEs: ", runtime.NumGoroutine())
	incr := 0
	var wg sync.WaitGroup
	var mutex sync.Mutex
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			mutex.Lock()
			tempVar := incr
			tempVar++
			incr = tempVar
			mutex.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(incr)
}
