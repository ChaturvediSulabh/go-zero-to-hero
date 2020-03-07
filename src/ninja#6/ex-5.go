// create a type SQUARE
// create a type CIRCLE
// attach a method to each that calculates AREA and returns it
// circle area= π r 2
// square area = L * W
// create a type SHAPE that defines an interface as anything that has the AREA method
// create a func INFO which takes type shape and then prints the area
// create a value of type square
// create a value of type circle
// use func info to print the area of square up to two places of decimal
// use func info to print the area of circle up to two places of decimal
package main

import (
	"fmt"
	"math"
)

type square struct {
	length  float64
	breadth float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.length * s.breadth
}

func (c circle) area() float64 {
	return math.Pi * c.radius
}

type shape interface {
	area() float64
}

func info(s shape) {
	area := s.area()
	fmt.Printf("%.2f\n", area)
}

func main() {
	s := square{
		length:  2.12,
		breadth: 2.12,
	}
	c := circle{
		radius: 4.0,
	}
	info(s)
	info(c)
}
