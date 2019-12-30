package main

import "fmt"

// Create a program that uses a switch statement with no switch expression specified.
func main() {
	switch {
	case false:
		fmt.Println("Won't Print")
	case false:
		fmt.Println("Won't Print")
	case true:
		fmt.Println("Will Print!")
	}
}
