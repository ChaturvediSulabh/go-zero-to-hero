// Building on the previous hands-on exercise, create a program that uses “else if” and “else”.

package main

import "fmt"

func main() {
	flag := false
	if flag {
		fmt.Println("ON")
	} else if !true {
		fmt.Println("OFF")
	} else {
		fmt.Println("OFF")
	}
}
