package main

import "fmt"

func main() {
	// 1.
	x := [5]int{10, 11, 12, 13, 14}

	// 2.
	for _, v := range x {
		fmt.Println(v)
	}

	// 3a.
	fmt.Printf("%T\n", x)
}
