package main

import "fmt"

func main() {
	// 1. assigns an int to a variable
	a := 52
	// 2. prints that int in decimal, binary, and hex
	fmt.Println(a)
	fmt.Printf("%d\n%b\n%x\n", a, a, a)
	// 3. shifts the bits of that int over 1 position to the left, and assigns that to a
	b := a << 1
	fmt.Printf("%d\n%b\n%x\n", b, b, b)

}
