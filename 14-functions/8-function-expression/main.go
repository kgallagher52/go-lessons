package main

import "fmt"

// Expression is something that comes down to a value

func main() {
	// func is like any other type
	f := func() {
		fmt.Println("My first func expression")
	}
	f()

	n := func(x int) {
		fmt.Println("My family #", x)
	}
	n(52)
}
