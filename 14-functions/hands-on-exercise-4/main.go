package main

import "fmt"

// 1. Create a user defined struct with  the identifier “person”
// the fields:
// first
// last
// age
type person struct {
	first string
	last  string
	age   int
}

//2 	attach a method to type person with
// the identifier “speak”
// the method should have the person say their name and age
func (p person) speak() {
	fmt.Println("Hi I am ", p.first, " ", p.last, "I am ", p.age, "years old!")
}

func main() {

	p1 := person{ // 3. create a value of type person

		first: "Keaton",
		last:  "Gallagher",
		age:   26,
	}
	p1.speak() // 4. call the method from the value of type person

}
