// Using iota, create 4 constants for the NEXT 4 years. Print the constant values.
package main

import "fmt"

const (
	_     = iota
	year1 = 2020 + iota
	year2 = 2020 + iota
	year3 = 2020 + iota
	year4 = 2020 + iota
)

func main() {
	fmt.Println(year1, year2, year3, year4)
}
