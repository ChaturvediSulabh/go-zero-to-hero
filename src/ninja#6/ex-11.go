// recursion in go with example
// Example: get nth number of fibonacci series
package main

import "fmt"

func main(){
	result := fib(8)
	fmt.Println(result)
}

func fib(n int)int{
	  if n == 0{
			return 0
		} else if n == 1{
			return 1
		} else {
		  return fib(n - 1) + fib(n-2)
		}
}