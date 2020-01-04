package main

import "fmt"

func main() {
	total := variadicSum(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("3.This is the total from the variadicSum ", total)
}

//** ... = means unlimited paramaters also known as lexical
// ...int = unlimited integers can be passed in

func variadicSum(x ...int) int {
	fmt.Println("1.", x)    // Print the value of the parameter passed in
	fmt.Printf("2.%T\n", x) // Print what type x is

	sum := 0

	for _, v := range x {
		sum += v // Adding every number of the numbers passed in to the variable sum
	}
	return sum
}
