package main

import "fmt"

type person struct {
	name    string
	element string
}

type bender struct {
	person
	good bool
}

// func (r reciever) identifier(parameters) (return(s)) { ... }

// Attaching method to a function * reciever goes in front of the identifier

func (s bender) speak() { // Defined reciver to speak so speak has access to bender which is a struct type
	fmt.Println("I am", s.name, s.element)
}

func main() {
	p1 := bender{
		person: person{
			"Katara",
			"Water",
		},
		good: true,
	}
	fmt.Println(p1)

	p2 := bender{
		person: person{
			"Ange",
			"Avatar",
		},
		good: true,
	}
	fmt.Println(p2)
	p1.speak() // Methods
	p2.speak() // Methods

}
