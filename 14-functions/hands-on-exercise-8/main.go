package main

import "fmt"

func main() {
	f := foo()
	fmt.Println(f())
}

// 1.Create a func which returns a func
func foo() func() int {
	return func() int {
		return 52
	}
}
