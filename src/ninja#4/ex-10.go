// Using the code from the previous example ex-9.go, delete a record from your map. Now print the map out using the “range” loop
package main

import "fmt"

func main() {
	dic := map[string][]string{
		"record1": []string{`bond_james`, `Shaken, not stirred`, `Martinis`, `Women`},
		"record2": []string{`moneypenny_miss`, `James Bond`, `Literature`, `Computer Science`},
		"record3": []string{`no_dr`, `Being evil`, `Ice cream`, `Sunsets`},
	}

	dic["record4"] = []string{"blah", "blah", "blah"}
	fmt.Println(dic)

	delete(dic, "record4")
	fmt.Println(dic)

	for _, v := range dic {
		fmt.Println(v)
		for i, j := range v {
			fmt.Println(i, j)
		}
	}
}
