package main

import "fmt"

const (
	a     = 43 // untyped constant
	b int = 43 // typed constant
)

func main() {
	fmt.Println(a, b)
}
