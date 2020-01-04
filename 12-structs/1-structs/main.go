package main

import "fmt"

type person struct { // ** Known as a aggregate data structure or complex data structure or composite data structure
	first string
	last  string
	age   int
}

func main() {
	// We created a type person
	p1 := person{ // * Allow you to aggrogate different values together
		first: "Harry", // Created the value of type person
		last:  "Potter",
		age:   24,
	}
	p2 := person{
		first: "Ron",
		last:  "Weasley",
		age:   26,
	}
	fmt.Println(p1)
	fmt.Println(p2)

	// Accessing fields from struct
	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2.first, p2.last, p2.age) // Composing together different data types

}
