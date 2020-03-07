// A “callback” is when we pass a func into a func as an argument. For this exercise,
// pass a func into a func as an argument
package main

import "fmt"

func bar(f func() string) string {
	return "Hello, this is callback"
}

func main() {
	g := func() string {
		return "success"
	}

	x := bar(g)
	fmt.Println(g())
	fmt.Println(x)
}
