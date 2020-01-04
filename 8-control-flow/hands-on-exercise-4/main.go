package main

import "fmt"

// Create a for loop using this syntax "for {}" have it print out the years you have been alive

func main() {
	year := 1993
	today := 2020
	for {
		if year <= today {
			fmt.Printf("%d\n", year)
			year++
		}
		if year > today {
			break
		}
	}
}
