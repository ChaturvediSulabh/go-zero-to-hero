// Create a map with a key of TYPE string which is a person’s “last_first” name, and a value of TYPE []string which stores their favorite things. Store three records in your map. Print out all of the values, along with their index position in the slice.
// 	`bond_james`, `Shaken, not stirred`, `Martinis`, `Women`
// `moneypenny_miss`, `James Bond`, `Literature`, `Computer Science`
// `no_dr`, `Being evil`, `Ice cream`, `Sunsets`
package main

import "fmt"

func main() {
	dic := map[string][]string{
		"record1": []string{`bond_james`, `Shaken, not stirred`, `Martinis`, `Women`},
		"record2": []string{`moneypenny_miss`, `James Bond`, `Literature`, `Computer Science`},
		"record3": []string{`no_dr`, `Being evil`, `Ice cream`, `Sunsets`},
	}

	for _, v := range dic {
		fmt.Println(v)
		for i, j := range v {
			fmt.Println(i, j)
		}
	}
}
