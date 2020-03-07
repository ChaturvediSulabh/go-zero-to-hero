// Use the “defer” keyword to show that a deferred func runs after the func containing it exits.
package main

import "fmt"

func main() {
	defer sum(1, 2)
	fmt.Println("Deferred function is printed below")
}

func sum(x int, y int) {
	fmt.Println(x + y)
}
