package main

import (
	"encoding/json"
	"fmt"
)

// Wizard ...
type Wizard struct {
	First string
	Age   int
}

func main() {
	w1 := Wizard{
		First: "Harry",
		Age:   21,
	}

	w2 := Wizard{
		First: "Hermine",
		Age:   22,
	}

	w3 := Wizard{
		First: "Ron",
		Age:   20,
	}

	wizards := []Wizard{w1, w2, w3}

	fmt.Println(wizards)

	mw, err := json.Marshal(wizards)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Marshaled Wizards: ", mw)

}
