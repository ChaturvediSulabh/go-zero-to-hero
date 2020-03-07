// Create a func which returns a func
// assign the returned func to a variable
// call the returned func
package main

import "fmt"

func foo() func() int {
	return func() int {
		return 42
	}
}

func main() {
	fmt.Println(foo()())
}
