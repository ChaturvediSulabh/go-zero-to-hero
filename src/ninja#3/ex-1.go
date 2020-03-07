// Print every number from 1 to 10,000
package main

import "fmt"

func main() {
	num := 1
	for num < 10001 {
		fmt.Println(num)
		num++
	}
}
