package main

import "fmt"

// Closure using a little code block to inclose certain code

func main() {
	a := incrementor()
	b := incrementor()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())

}

func incrementor() func() int { //return a function
	var x int // Value will be zero
	return func() int {
		x++
		return x
	}
}

/*  is inclosed around x
func() int {
	x++
	return x
}

*/
