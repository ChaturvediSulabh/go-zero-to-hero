// Create and use an anonymous struct.
package main

import (
	"fmt"
	"strings"
)

func main() {
	p := struct {
		name    string
		contact map[string]int
		address []string
	}{
		name: "Sulabh Chaturvedi",
		contact: map[string]int{
			"home":   222222,
			"office": 111111,
		},
		address: []string{
			"flat xx",
			"yy road",
			"London",
		},
	}

	fmt.Printf("Name: %v\nHome: %v\nOffice: %v\nAddress: %v\n", p.name, p.contact["home"], p.contact["office"], strings.Join(p.address[:], " "))
}
