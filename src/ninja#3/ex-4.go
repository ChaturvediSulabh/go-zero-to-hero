// Hands-on exercise #4
// Create a for loop using this syntax
// for { }
// Have it print out the years you have been alive.
package main

import "fmt"

func main() {
	i := 1986
	for {
		if i == 2020 {
			break
		} else {
			fmt.Println(i)
		}
		i++
	}
}
