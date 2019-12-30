package main

import "fmt"

func main() {
	var x [5]int
	fmt.Println(x)
	x[3] = 52 // This would assign the fourth position value in the array to 52 because we start counting at 0
	fmt.Println(x)
	fmt.Println(len(x)) // Getting the length of an array

}
