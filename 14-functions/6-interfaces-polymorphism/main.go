package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person // taking in type person which is an underlyning struct
	ltk    bool
}

// Method
func (s secretAgent) speak() {
	fmt.Println(s)
	// fmt.Println("I am", s.first, s.last)
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last, "- the person speak")
}

// ** A VALUE CAN BE MORE THEN ONE TYPE

// ** Interfaces allows us to define behavior and then also alow us to do polymorphism
type human interface {
	speak()
}

func bar(h human) {
	//Insertion
	switch h.(type) { // Special case of switch statment when you can switch on a type
	case person:
		fmt.Println("INSERTION - I was passed into bar", h.(person).first)

	case secretAgent:
		fmt.Println("INSERTION - I was passed into bar", h.(secretAgent).first)

	}
	fmt.Println("I was passed into bar", h)
}

type hotdog int

func main() {
	sa1 := secretAgent{
		person: person{
			"james",
			"Bond",
		},
		ltk: true,
	}

	sa2 := secretAgent{
		person: person{
			"Riddler",
			"Tiddler",
		},
		ltk: false,
	}

	p1 := person{
		first: "Dr.",
		last:  "Yes",
	}
	fmt.Println(sa1)
	sa1.speak()
	sa2.speak()

	fmt.Println(p1)

	bar(sa1)
	bar(sa2)
	bar(p1)

	// conversion
	var x hotdog = 52
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	var y int
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)

}
