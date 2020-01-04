package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 67, 8, 9}
	//	s := sum(nums) this will not work because the function want's an int not a slice of int's
	s := sum(nums...)
	fmt.Println("Sum of all numbers", s)

	// 3
	s2 := even(sum, nums...)
	fmt.Println("Sum of all even numbers", s2)

	// 4
	s3 := odd(sum, nums...)
	fmt.Println("Sum of all odd numbers", s3)
}

func sum(xi ...int) int { // Sum has a variadic argument underlyning int that returns an int
	total := 0

	for _, v := range xi { // Ranging over the integers passed in and adding them up to the total
		total += v
	}
	return total
}

// Want all even numbers
// func even(f func(xi ...int) int, vi ...int) int {
// f func(xi ...int) int, - is the call back function | vi ...int - the second parameter being passed in
// 3
func even(f func(xi ...int) int, vi ...int) int { // Sum of all even numbers
	var yi []int // Slice to hold values
	for _, v := range vi {
		if v%2 == 0 { // Checking if they are even
			yi = append(yi, v) // appending the evens to the slice we created above
		}
	}
	return f(yi...) // Returning
}

// Creating odd numbers
// 3 odd
func odd(f func(xi ...int) int, vi ...int) int { // Sum of all even numbers
	var yi []int // Slice to hold values
	for _, v := range vi {
		if v%2 != 0 { // Checking if they are even
			yi = append(yi, v) // appending the evens to the slice we created above
		}
	}
	return f(yi...) // Returning
}
