package main

import "fmt"

type person struct {
	first      string
	last       string
	favFlavors []string
}

func main() {
	p1 := person{
		first: "Keaton",
		last:  "Gallagher",
		favFlavors: []string{
			"Mint",
			"vanilla",
		},
	}
	// 1. Take the code from the previous exercise, then store the values of type person in a map with the key of last name. Access each value in the map. Print out the values, ranging over the slice.

	m := map[string]person{ // p1.last = key | p1 = value you want for the record found
		p1.last: p1,
	}
	fmt.Printf("%T\n", m)

	fmt.Println("Map of last name value\n", m)

}
