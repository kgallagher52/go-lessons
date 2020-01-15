package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First   string // ** These have to be capital if you plan on using them in a brought in package
	Element string
	Age     int
}

func main() {
	p1 := person{
		First:   "Aang",
		Element: "All",
		Age:     100,
	}
	p2 := person{
		First:   "Katara",
		Element: "Water",
		Age:     24,
	}

	people := []person{p1, p2}

	fmt.Println("Before Marshal: ", people)

	// byte slice Marshaling
	bs, err := json.Marshal(people)
	if err != nil { // Check for errors on line 30
		fmt.Println("Error while marshaling: ", err)
	}
	// Convert our byte slice to string if no error
	fmt.Println("Marshaled Value: ", string(bs))

}
