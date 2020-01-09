package main

import "math"

import "fmt"

type shape interface {
	area() float64
}

type circle struct {
	radius float64
}

// Functions with method area
// Non poi
func (c *circle) area() float64 { // Is of type shape
	return math.Pi * c.radius * c.radius
}

// Function

func info(s shape) {
	fmt.Println("area", s.area())
}

func main() {
	c := circle{9}
	fmt.Printf("%T\n", &c)
	info(&c)
}
