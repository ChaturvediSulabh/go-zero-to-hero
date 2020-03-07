// Create a program that uses a switch statement with the switch expression specified as a variable of TYPE string with the IDENTIFIER “favSport”.
package main

import "fmt"

var favSport string

func main() {
	switch favSport {
	case "Soccer":
		fmt.Println("Soccer")
	default:
		fmt.Println("Cricket")
	}
}
