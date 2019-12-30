package main

import "fmt"

func main() {
	x := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(x)
	fmt.Println(x[1])
	fmt.Println(x[3:])
	fmt.Println(x[1:3]) // * We will go up to 3 but not including it

	// Using a range clause
	for i, v := range x {
		fmt.Println(i, v)
	}

	// * Using Slices in a for loop
	for i := 0; i <= 4; i++ {
		fmt.Println(i, x[i])
	}

}
