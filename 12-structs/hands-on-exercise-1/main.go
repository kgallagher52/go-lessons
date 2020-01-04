package main

import "fmt"

// 1. Creat your own type which has an underlying type of "struct"
type person struct {
	first      string
	last       string
	favFlavors []string
}

func main() {
	// 2. Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice which stores the favorite flavors.

	p1 := person{
		first: "Keaton",
		last:  "Gallagher",
		favFlavors: []string{
			"Mint",
			"vanilla",
		},
	}

	p2 := person{
		first: "Ali",
		last:  "Witt",
		favFlavors: []string{
			"CookieDough",
			"vanilla",
		},
	}

	fmt.Println("1. Person first name\n\t", p1.first)
	fmt.Println("1. Person last name\n\t", p1.last)
	fmt.Println("Flavors:")

	for _, v := range p1.favFlavors { // Ranging the flavors
		fmt.Println("\t", v)
	}
	fmt.Println("2. Person first name\n\t", p2.first)
	fmt.Println("2. Person last name\n\t", p2.last)
	fmt.Println("Flavors:")

	for _, v := range p2.favFlavors { // Ranging the flavors
		fmt.Println("\t", v)
	}

}
