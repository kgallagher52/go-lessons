package main

import "fmt"

// Create a program that uses a switch statement with the switch expression specified as a variable of TYPE string with the IDENTIFIER “favSport"

func main() {
	favSport := "Snowboarding"

	switch favSport {
	case "Skating":
		fmt.Println("Negative")
	case "Golf":
		fmt.Println("Negative")
	case "Snowboarding":
		fmt.Println("Postive")
	case "Football":
		fmt.Println("Negative")
	}
}
