// Create your own type “person” which will have an underlying type of “struct” so that it can store the following data:
// first name
// last name
// favorite ice cream flavors
// Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice which stores the favorite flavors.
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
		lastName:  "Chaturvedi",
		iceCreamFlavours: []string{
			"plain vanilla",
			"tender coconut",
		},
	}

	for i, v := range p1.iceCreamFlavours {
		fmt.Println("Sulabh's favourite flavours:", i+1, "-", v)
	}

	for i, v := range p2.iceCreamFlavours {
		fmt.Println("Seema's favourite flavours:", i+1, "-", v)
	}
}
