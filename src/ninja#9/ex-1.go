// in addition to the main goroutine, launch two additional goroutines
// each additional goroutine should print something out
// use waitgroups to make sure each goroutine finishes before your program exists
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("Total cpu's: ", runtime.NumCgoCall())
	counter := 2

	var wg sync.WaitGroup
	wg.Add(counter)

	for i := 0; i < counter; i++ {
		go func() {
			fmt.Println("GOROUTINE", runtime.NumGoroutine())
			wg.Done()
		}()
	}
	wg.Wait()
}
