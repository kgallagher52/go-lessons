package main

import "fmt"

func main() {
	a := 42 // 1. Assign a to 42 which is an int
	// & = operator that shows you the address in memory
	fmt.Println("Memory in address", &a)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\nPointer too an int", &a)

	b := &a // 2. Assign b to a's address
	fmt.Println("Value passed from a : b = ", b)
	// *b - the star here is an operator
	// *int - is different it's a type
	// * - Operator derefrence the address
	fmt.Println("Derefrenced address : *b = ", *b)
	fmt.Println("Give me the address then derefrence it ", *&a)

	*b = 43 // 3. Give me the value of b and assign it to 43
	fmt.Println(a)
}
