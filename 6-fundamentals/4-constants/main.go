package main

import "fmt"

const a = 42
const b = 42.78
const c = "Harry Potter"

// Another way of writing this would be:

const (
	d int     = 42
	e float64 = 42.78
	f string  = "Harry Potter"
)

func main() {
	fmt.Println(a, d)
	fmt.Println(b, e)
	fmt.Println(c, f)
	fmt.Printf("%T\n%T\n", a, d)
	fmt.Printf("%T\n%T\n", b, e)
	fmt.Printf("%T\n%T\n", c, f)

}
