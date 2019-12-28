package main

import "fmt"

// 1.
type hotDog int

var x hotDog
var y int

func main() {
	// 2a1. print out the value of the variable “x”
	fmt.Println(x)
	// 2a2. print out the type of the variable “x”
	fmt.Printf("%T\n", x)
	// 2a3. assign your own VALUE to the VARIABLE “x” using the “=” OPERATOR
	x = 52
	// 2a4. print out the value of the variable “x”
	fmt.Println(x)
	// 2b1. convert the TYPE of the VALUE stored in “x” to the UNDERLYING TYPE
	// 2b1a. then use the “=” operator to ASSIGN that value to “y”
	y = int(x)
	// 2b1b. print out the value stored in “y”
	fmt.Println(y)
	// 2b1c. print out the type of “y”
	fmt.Printf("%T\n", y)

}
