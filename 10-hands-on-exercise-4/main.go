package main

import "fmt"

func main() {
	// 1
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	// 2
	x = append(x, 52)
	// 3
	fmt.Println(x)
	// 4
	x = append(x, 53, 54, 55)
	// 5
	fmt.Println(x)
	// 6
	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	// 7
	fmt.Println(x)

}
