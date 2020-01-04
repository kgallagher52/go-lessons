package main

import "fmt"

func main() {
	defer foo() // defer makes that piece run on close or exit
	bar()
}

func foo() {
	fmt.Println("1. Foo fired")
}

func bar() {
	fmt.Println("2. Bar fired")
}
