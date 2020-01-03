package main

import "fmt"

// 1. Create and use an anonymous struct.

func main() {
	wizard := struct {
		name  string
		color string
		age   int
	}{
		name:  "Gandoff",
		color: "Grey",
		age:   200,
	}
	fmt.Println("Anonymous Struct: ", wizard)
	fmt.Printf("%T\n", wizard)
}
