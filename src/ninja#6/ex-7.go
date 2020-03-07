// Assign a func to a variable, then call that func
package main

import "fmt"

func main() {
	x := func() string {
		return "Hello, XAnonymous"
	}
	fmt.Println(x())
}
