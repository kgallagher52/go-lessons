package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	// ** These have to be capital if you plan on using them in a brought in package
	First   string `json:"First"`
	Element string `json:"Element"`
	Age     int    `json:"Age"`
}

func main() {
	js := `[{"First":"Aang","Element":"All","Age":100},{"First":"Katara","Element":"Water","Age":24}]`

	// Convert to slice of bytes
	bs := []byte(js)
	fmt.Println("Slice of bytes", bs)
	fmt.Printf("Type bs:  %T \n", bs)
	fmt.Printf("Type js:   %T \n", js)

	// var people []person - this is the same as below
	people := []person{}

	if err := json.Unmarshal(bs, &people); err != nil {
		fmt.Println(err)
	}
	fmt.Println("All of the Data: ", people)

	for _, p := range people {
		fmt.Println("Person: ", p.First, p.Element, p.Age)
	}
}
