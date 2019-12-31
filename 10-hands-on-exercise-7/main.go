package main

import "fmt"

func main() {
	// 1
	fs := []string{"James", "Bond", "Shaken, not stirred"}
	ss := []string{"Miss", "Moneypenny", "Hellooooo, James"}
	fmt.Println(fs)
	fmt.Println(ss)

	ms := [][]string{fs, ss}
	fmt.Println(ms)

	for _, v := range ms {
		fmt.Println(v)
		for _, sv := range v {
			fmt.Println(sv)
		}
	}
}
