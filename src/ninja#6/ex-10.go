// Closure is when we have “enclosed” the scope of a variable in some code block. For this hands-on exercise, create a func which “encloses” the scope of a variable:
package main

import "fmt"

func closure(x string) string {
	return "Scope of x is within the func closure"
}

func main() {
	x := closure("variable") // this variable x has scope of func main and is different from the parameter in func closure
	fmt.Println(x)
}
