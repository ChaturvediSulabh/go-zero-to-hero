// write a program that
// launches 10 goroutines
// each goroutine adds 10 numbers to a channel
// pull the numbers off the channel and print them

package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	defer receive(c)
	for i := 0; i < 10; i++ {
		send(c)
	}
}

func send(c chan int) <-chan int {
	go func(c chan<- int) {
		for j := 0; j < 10; j++ {
			c <- j
		}
		close(c)
	}(c)
	return c
}
func receive(c <-chan int) {
	for n := range c {
		fmt.Println(n)
	}
}
