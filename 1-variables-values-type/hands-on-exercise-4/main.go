package main

import "fmt"

// 1. Create custom type
type customType int

// 2. create a VARIABLE of your new TYPE with the IDENTIFIER “x” using the “VAR” keyword
var x customType

func main() {
	// 3a. print out the value of the variable “x”
	fmt.Println(x)
	// 3b. print out the type of the variable “x”
	fmt.Printf("%T\n", x)
	// 3c. assign 42 to the VARIABLE “x” using the “=” OPERATOR
	x = 42
	// 3d. print out the value of the variable “x”
	fmt.Println(x)

}
