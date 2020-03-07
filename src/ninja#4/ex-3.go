// Using the code from the previous example, use SLICING to create the following new slices which are then printed:
// [42 43 44 45 46]
// [47 48 49 50 51]
// [44 45 46 47 48]
// [43 44 45 46 47]
package main

import "fmt"

func main() {
	sl := []int{42, 43, 44, 45, 46}
	fmt.Println(sl)

	sl = append(sl, 47, 48, 49, 50, 51)

	fmt.Println(sl[5:])
	fmt.Println(sl[2:7])
	fmt.Println(sl[1:6])
}
