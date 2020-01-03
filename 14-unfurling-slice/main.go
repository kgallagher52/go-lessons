package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// sum function is expecting unlimited number of int's ...int since we are passing in a slice of int;s we need to do xi... so the program knows we want to dump it
	total := sum(xi...)
	fmt.Println(total)
}

func sum(xi ...int) int {
	sum := 0

	for _, v := range xi {
		sum += v
	}
	return sum
}
