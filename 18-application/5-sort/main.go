package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{10, 0, 1, 100, 52, 14, 37}
	xs := []string{"Katara", "Zuko", "Toph", "Sakka", "Momo", "Aang"}

	fmt.Println(xi)
	fmt.Println(xs)

	// ** the sort function takes in (a []int) but does not return so we do not need to assign paramater to this
	sort.Ints(xi)
	fmt.Println(xi)

	sort.Strings(xs)
	fmt.Println(xs)

}
