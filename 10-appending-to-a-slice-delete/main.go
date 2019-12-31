package main

import "fmt"

func main() {
	x := []int{1, 2, 3}
	fmt.Println("Original Slice: ", x)
	x = append(x, 4, 5, 6)
	fmt.Println("Appended Slice: ", x)

	y := []int{7, 8, 9}
	x = append(x, y...) //The y... is just saying take all the values and append it after all of the values in the x slice
	fmt.Println("Appended with another Slice: ", x)

	x = append(x[:2], x[4:]...) // Deleting up to index 2 till index 4 but not including 4 results: [1,2,5,6,7,8,9]
	fmt.Println("Deleting from slice: ", x)

}
