// Using the following operators, write expressions and assign their values to variables:
// ==
// <=
// >=
// !=
// <
// >
// Now print each of the variables.
package main

import "fmt"

const (
	a = iota
	b
)

func main() {
	c := (a == b)
	d := (a <= b)
	e := (a >= b)
	f := (a != b)
	g := (a > b)
	h := (a < b)
	fmt.Printf("%v\n%v\n%v\n%v\n%v\n%v\n", c, d, e, f, g, h)
}
