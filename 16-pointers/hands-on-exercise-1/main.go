package main

import "fmt"

func main() {
	// 1. Create a value and assign it to a variable.
	x := 45
	// 2. Print the address of that value.
	fmt.Println("Address of value x: \n", &x)
}
