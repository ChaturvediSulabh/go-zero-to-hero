// Write a program that
// assigns an int to a variable
// prints that int in decimal, binary, and hex
// shifts the bits of that int over 1 position to the left, and assigns that to a variable
// prints that variable in decimal, binary, and hex
package main

import "fmt"

func main() {
	i := 4
	fmt.Printf("Decimal -> %v\tBinary -> %b\tHexadecimal -> %#x\n", i, i, i)
	j := i << 1
	fmt.Printf("Decimal -> %v\tBinary -> %b\tHexadecimal -> %#x", j, j, j)
}
