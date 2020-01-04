package main

import "fmt"

func main() {
	// 1
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	// 2
	for i, v := range x {
		fmt.Println(i, v)
	}

	//3
	fmt.Printf("%T\n", x)
}
