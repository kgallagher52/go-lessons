package main

import "fmt"

func main() {
	// 1b. pass in a value of type []int into your func (unfurl the []int)
	n := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sum := foo(n...)
	fmt.Println(sum)

	n2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sum2 := bar(n2)
	fmt.Println(sum2)

}

// 1. create a func with the identifier foo that
func foo(v ...int) int { // 1a. Pass in a variadic parameter of type int
	//1c. returns the sum of all values of type in passed in
	total := 0
	for _, v := range v {
		total += v
	}
	return total
}

// 2. Create a func with the identifier bar
func bar(s []int) int { // 2a. takes in a parameter of type []int
	total := 0
	// 2b. returns the sum of all values of type int passed in
	for _, s := range s {
		total += s
	}
	return total
}
