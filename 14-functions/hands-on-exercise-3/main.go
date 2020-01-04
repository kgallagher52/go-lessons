package main

import "fmt"

func main() {
	first()
	second()
	// 1. Use the “defer” keyword to show that a deferred func runs after the func containing it exits.
	defer first()
	second()
}

func first() {
	fmt.Println("This function should run first!!!")
}
func second() {
	fmt.Println("This function should run second!!!")
}
