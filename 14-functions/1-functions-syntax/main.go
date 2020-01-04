package main

import "fmt"

func main() {
	foo()

	bar("Cheese Biscuts")

	S1 := hoo("raw")
	fmt.Println(S1)

	n, b := element("What nation?", "Good or Bad?")
	fmt.Println("Nation:", n)
	fmt.Println("Good?", b)

}

// func (r receiver) identifier(paramaters) (return(s)) { ... }

// No paramater function
func foo() {
	fmt.Println("Hello from foo")
}

// Takes an argument function
// * EVERYTHING IN GO IS PASSED BY VALUE
func bar(s string) {
	fmt.Println("Hello from", s)
}

// Function takes an argument and returns string
func hoo(s string) string {
	return fmt.Sprint("Hello from hoo", s)
}

// Function takes in two arguments and returns a string and bool
func element(el string, gb string) (string, bool) {
	fmt.Println(el)
	a := fmt.Sprint("FIRE")

	fmt.Println(gb)
	b := false

	return a, b

}
