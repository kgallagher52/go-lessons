package main

import "fmt"

// 1. Closure is when we have “enclosed” the scope of a variable in some code block. For this hands-on exercise, create a func which “encloses” the scope of a variable:

func main() {
	fmt.Println("I am about to close this variable")
	{
		x := "String"
		fmt.Println("I am about to close this variable", x)

	}

}
