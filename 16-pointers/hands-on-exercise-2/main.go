package main

import "fmt"

type person struct { // 1. create a person struct
	name string
}

func main() {
	p1 := person{
		name: "Harry Potter",
	}
	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)
}

func changeMe(p *person) { // 2. create a func called “changeMe” with a *person as a parameter
	p.name = "Mr.Tom" // 2a. change the value stored at the *person address
	// (*p).name = "Tommy"
}
