package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type wizard struct {
	First  string
	Last   string
	Age    int
	Powers []string
}

func main() {
	w1 := wizard{
		First: "Harry",
		Last:  "Potter",
		Age:   21,
		Powers: []string{
			"Blade Shooting",
			"Dragon Summon",
			"Invisiblity",
		},
	}

	w2 := wizard{
		First: "Hermine",
		Last:  "Granger",
		Age:   22,
		Powers: []string{
			"Repair Anything",
			"Extract Information",
			"Erase Memories",
		},
	}

	w3 := wizard{
		First: "Ron",
		Last:  "Weasly",
		Age:   20,
		Powers: []string{
			"Light Fires",
			"Win Any Game",
			"Breath Underwater",
		},
	}

	wizards := []wizard{w1, w2, w3}

	fmt.Println("\n", wizards)

	if err := json.NewEncoder(os.Stdout).Encode(wizards); err != nil {
		fmt.Println(err)
	}
}
