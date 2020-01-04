package main

import "fmt"

func main() {
	// 3&4. Call both functions and print them out
	fmt.Println(foo())
	fmt.Println(bar())
}

// 1. create a func with the identifier foo that returns an int
func foo() int {
	return 3
}

// 2. create a func with the identifier bar that returns an int and a string
func bar() (int, string) {
	return 52, "Avatar Time"
}
