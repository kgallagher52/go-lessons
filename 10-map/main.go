package main

import "fmt"

func main() {
	m := map[string]int{ // Set key to string type & the value to an integer type
		"Harry":  21,
		"Hagrid": 180,
	}
	fmt.Println(m)
	// Accessing the data by key
	fmt.Println("Record that exists: ", m["Harry"])

	fmt.Println("Record that does not exist: ", m["Lily"]) // When you enter a key that is not in your map it will return the zero value

	v, ok := m["Lily"] // Checking if the record does exist using the ok ideom
	fmt.Println("Value returned from record check: ", v)
	fmt.Println("Boolean on if the record exists: ", ok)

	// Adding to map
	m["Lily"] = 21
	fmt.Println("Map after adding record to it: ", m)

	if v, ok := m["Harry"]; ok { // * Very common used ideom in Golang for testing records
		fmt.Println("THIS IS THE IF PRINT", v)
		// This would not print if m["Lily"] because it would return as false
	}

	// Deleting from map
	delete(m, "Harry")

	// Mapping over a map
	fmt.Println("Mapped out map:")
	for k, v := range m {
		fmt.Println("\tIndex", k, "\tValue", v)
	}

	// Mapping over a slice
	fmt.Println("Mapped out Slice:")
	xi := []int{1, 2, 3, 4, 5}
	for i, x := range xi {
		fmt.Println("\tIndex:", i, "\tValue", x)

	}

}
