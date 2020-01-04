package main

import "fmt"

func main() {
	// 1: Returning a string
	s1 := foo()
	fmt.Println(s1)

	// 1b: Returning cleanre way of string
	s2 := foo2()
	fmt.Println(s2)

	// 2: Func int
	x := bar()
	fmt.Printf("%T\n", x) // Printing the type of func int

	// 2b: We can also execute func int
	fmt.Println(x())

	// 2c: Cleaner way of writing 2 and 2b together
	fmt.Println("runs bar which is a function being returned then runs the returning function", bar()())
}

// 1: Returning a string
func foo() string {
	s := "Returning foo string"
	return s
}

// 1b: Cleaner way of writing foo
func foo2() string {
	return "Returning claner code string"
}

// 2: Function returning a function
func bar() func() int {
	return func() int { //Anonymous function right here
		return 52
	}
}
