package main

import "fmt"

var a int

type wizardAge int // Creating custom type

var b wizardAge // Declaring b to custom type

func main() {
	a = 52
	b = 200 // Assigning value to b which is a custom type = int
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
}
