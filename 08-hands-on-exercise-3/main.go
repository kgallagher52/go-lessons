package main

import "fmt"

// Create a for loop with this syntax "for condition {}" Have it print out the years you have been alive
func main() {
	birth := 1993
	today := 2020
	for birth <= today {
		fmt.Printf("%d\n", birth)
		birth++
	}
}
