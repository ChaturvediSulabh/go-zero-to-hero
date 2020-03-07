// Create TYPED and UNTYPED constants. Print the values of the constants.
package main

import "fmt"

const (
	a int     = 0
	b float64 = 42.44
	d string  = "otherthing"
)

func main() {
	f := 0
	g := "something"

	fmt.Println(a, b, d, f, g)
}
