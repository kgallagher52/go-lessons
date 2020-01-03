package main

import "fmt"

// 1. Create and use an anonymous struct.

func main() {
	wizard := struct {
		name   string
		color  string
		age    int
		pets   []string
		spells map[string]int
	}{
		name:  "Gandoff",
		color: "Grey",
		age:   200,
		pets:  []string{"dog", "cat", "owl"},
		spells: map[string]int{
			"Fireball":         10,
			"Lightning Strike": 100,
			"nail growing":     1,
		},
	}

	fmt.Println("Anonymous Struct: ", wizard)
	fmt.Printf("%T\n", wizard)
}
