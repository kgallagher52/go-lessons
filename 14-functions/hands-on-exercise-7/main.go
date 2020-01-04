package main

import (
	"fmt"
)

var x int
var g func() // Defined the type as a function

func main() {

	f := func() { // Made f identifer which is type function
		for i := 0; i < 3; i++ {
			fmt.Println(i)
		}
	}

	f() // Calling f which is underlyn type func
	fmt.Printf("%T\n", f)

	fmt.Println(x)
	fmt.Printf("%T\n", x)

	g = f
	g()
	fmt.Printf("this is g %T\n", g)

	fmt.Println("done")
}
