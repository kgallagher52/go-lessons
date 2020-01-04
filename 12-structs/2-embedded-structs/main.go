package main

import "fmt"

type person struct { // ** Known as a aggregate data structure or complex data structure or composite data structure
	first string // * Have to be unique
	last  string
	age   int
}

// Embedding Structs
type orderPheonix struct {
	person // Include everything that a person offers
	wizard bool
}

// You can declare variables that are the same type in the same field
var y, x int

func main() {
	y = 12
	x = 10
	fmt.Println(y, x)
	oP1 := orderPheonix{ // Embedding structs part two
		person: person{ // Gets promoted to the outer type of order Pheonix
			first: "Herminie",
			last:  "Granger",
			age:   23,
		},
		wizard: true,
	}
	p2 := person{
		first: "Ron",
		last:  "Weasley",
		age:   26,
	}

	fmt.Println(oP1)
	fmt.Println(p2)

	// Accessing fields from struct
	fmt.Println(oP1.first, oP1.last, oP1.age)
	fmt.Println(p2.first, p2.last, p2.age) // Composing together different data types

}
