package main

import (
	"fmt"
	"math"
)

// 1, create a type SQUARE
type square struct {
	length float64
}

// 2. create a type CIRCLE
type circle struct {
	radius float64
}

// 3. attach a method to each that calculates AREA and returns it
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (s square) area() float64 {
	return s.length * s.length
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	circ := circle{
		radius: 12.345,
	}
	squa := square{
		length: 15,
	}
	info(circ)
	info(squa)
}
