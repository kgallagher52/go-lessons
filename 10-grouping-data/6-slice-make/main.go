package main

import "fmt"

func main() {
	// If you know the specific size of an array you need to use then thats when you can use make

	x := make([]int, 10, 1000)
	fmt.Println("make Slice: ", x)
	fmt.Println("Length ", len(x))
	fmt.Println("Capacity", cap(x))

}
