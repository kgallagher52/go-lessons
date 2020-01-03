package main

import "fmt"

// 1. Create a new type vehical
type vehical struct { // 1a. The underlying type is a struct
	doors int
	color string
}

// 2. Create two new types truck & sedan
type truck struct { // 2a. The underlying type is a struct
	vehical        // 2b. Embeded is the "vehicle" type
	fourWheel bool // 2c. Give truck the field "fourWheel" which will be set to bool

}

// 2. Create two new types truck & sedan
type sedan struct { // 2a. The underlying type is a struct
	vehical      // 2b. Embeded is the "vehicle" type
	luxury  bool // 2d. Give sedan the field "luxury" which will  be set to bool
}

func main() {
	// 3b. using a composite literal, create a value of type truck and assign values to the fields;
	t := truck{
		vehical: vehical{
			doors: 4,
			color: "Red",
		},
		fourWheel: true,
	}

	// 3c. using a composite literal, create a value of type sedan and assign values to the fields.
	s := sedan{
		vehical: vehical{
			doors: 2,
			color: "Silver",
		},
		luxury: false,
	}

	fmt.Println(t) // 4. Print out each of these values.
	fmt.Println(s) // 4. Print out each of these values.

	fmt.Println(t.fourWheel) // 5. Print out a single field from each of these values.
	fmt.Println(s.doors)     // 5. Print out a single field from each of these values.

}
