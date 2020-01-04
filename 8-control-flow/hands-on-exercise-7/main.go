package main

import "fmt"

// Build of previous exercise, create a program that uses "else if" and "else"
func main() {
	x := "Harry Potter"
	if x != "Harry Potter" {
		fmt.Println(x)
	} else if x == "Harry Potter" {
		fmt.Println("if else", x)
	} else {
		fmt.Println("X does not equal the conditions")
	}
}
