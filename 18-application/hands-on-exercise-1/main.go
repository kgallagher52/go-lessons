package main

import (
	"fmt"
)

type wizard struct {
	First string
	Age   int
}

func main() {
	w1 := wizard{
		First: "Harry",
		Age:   21,
	}

	w2 := wizard{
		First: "Hermine",
		Age:   22,
	}

	w3 := wizard{
		First: "Ron",
		Age:   20,
	}

	wizards := []wizard{w1, w2, w3}

	fmt.Println(wizards)

	// your code goes here
}
