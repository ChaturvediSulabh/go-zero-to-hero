// Take the code from the previous exercise, then store the values of type person in a map with the key of last name. Access each value in the map. Print out the values, ranging over the slice.
package main

import "fmt"

type person struct {
	firstName        string
	lastName         string
	iceCreamFlavours []string
}

func main() {
	p1 := person{
		firstName: "Sulabh",
		lastName:  "Chaturvedi",
		iceCreamFlavours: []string{
			"tender coconut",
			"natural choclate",
		},
	}

	p2 := person{
		firstName: "Seema",
		lastName:  "Sharma",
		iceCreamFlavours: []string{
			"plain vanilla",
			"tender coconut",
		},
	}

	m := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}

	for _, v := range m {
		for _, fl := range v.iceCreamFlavours {
			fmt.Println(v.firstName+" "+v.lastName, "likes", fl)
		}
	}
}
