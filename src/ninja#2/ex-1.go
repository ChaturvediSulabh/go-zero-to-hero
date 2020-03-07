// Write a program that prints a number in decimal, binary, and hex
package main

import "fmt"

var x int

func main() {
	fmt.Println("Decimal ->", x)
	x = 9
	fmt.Printf("Binary ->%b\n", x)
	fmt.Printf("Hexadecimal ->%#x", x)
}
