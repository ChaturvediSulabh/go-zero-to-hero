// Hands-on exercise #1
// Using the short declaration operator, ASSIGN these VALUES to VARIABLES with the IDENTIFIERS “x” and “y” and “z”
// 42
// “James Bond”
// true
// Now print the values stored in those variables using
// a single print statement
// multiple print statements

package main

import "fmt"

func main() {
	x := 42
	y := "James Bond"
	z := true

	fmt.Printf("x = %v\ty = %v\tz = %v\n", x, y, z)
	fmt.Println("x =", x)
	fmt.Println("y =", y)
	fmt.Println("z =", z)
}
