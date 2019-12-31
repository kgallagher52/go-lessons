package main

import "fmt"

func main() {
	// 1
	m := map[string][]string{ // * ** Complex Mapping
		"Harry_Potter":    []string{"Hagrid", "Alavanders", "Ron"},
		"Keaton_Gallager": []string{"Snowboarding", "GF", "Movies"},
		"James Franco":    []string{"Acting", "Lauging", "Money"},
	}
	fmt.Println(m)

	m["Jeremy Ranck"] = []string{"Gym", "Overwatch", "Utah Utes"}
	fmt.Println("Record added", m)

	for k, v := range m {
		fmt.Println("This is the record for", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}

}
