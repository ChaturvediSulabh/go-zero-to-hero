// Create a slice of a slice of string ([][]string). Store the following data in the multi-dimensional slice:

// "James", "Bond", "Shaken, not stirred"
// "Miss", "Moneypenny", "Helloooooo, James."

// Range over the records, then range over the data in each record.
package main

import (
	"fmt"
)

func main() {
	sliceOfSlice := [][]string{}

	sl1 := []string{"James", "Bond", "Shaken, not stirred"}
	sl2 := []string{"Miss", "Moneypenny", "Helloooooo, James."}

	sliceOfSlice = append(sliceOfSlice, sl1, sl2)
	fmt.Println(sliceOfSlice)

	for i, v := range sliceOfSlice {
		for _, x := range sliceOfSlice[i] {
			fmt.Println(x, v)
		}
	}
}
