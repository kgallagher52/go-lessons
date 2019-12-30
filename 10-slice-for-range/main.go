package main

import "fmt"

func main() {
	x := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(len(x)) // Getting length of my slice of integers
	fmt.Println(x[0])   // Getting specific position
	fmt.Println(x[1])   // Getting specific position
	fmt.Println(x[2])   // Getting specific position
	fmt.Println(x[3])   // Getting specific position
	fmt.Println(x[4])   // Getting specific position
	fmt.Println(x[5])   // Getting specific position

	for i, v := range x { // i = index | v = value
		fmt.Println(i, v)
	}

}
