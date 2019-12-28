package main

import "fmt"

var x int
var y string
var z bool

func main() {
	// 1. Assign values to variables
	x = 42
	y = "James Bond"
	z = true
	fmt.Println(x, y, z)
	// 2. Second part
	s := fmt.Sprintf("%v\t%v\t%v\t", x, y, z)
	fmt.Println(s)
}
