// write a program that
// puts 100 numbers to a channel
// pull the numbers off the channel and print them

package main

import "fmt"

func main() {
	var numbers <-chan int
	numbers = send()
	receive(numbers)
}

func send() <-chan int {
	c := make(chan int)
	go func() {
		for i := 1; i < 101; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}

func receive(numbers <-chan int) {
	for n := range numbers {
		fmt.Println(n)
	}
}
