package main

import "fmt"

func main() {
	// 1
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x)
	// 2
	x = append(x[:3], x[6:]...)
	fmt.Println(x)

}
