// Create a program that uses a switch statement with no switch expression specified

package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("OFF")
	case true:
		fmt.Println("ON")
		fallthrough
	default:
		fmt.Println("ERR")
	}
}
