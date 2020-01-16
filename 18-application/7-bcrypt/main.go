package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	p := `mySecretPassword`
	fmt.Println("Not hashed Pass:", p)

	hp, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Hashed Pass:", hp)

	cp := "mySecretPassword"

	// Example of it working
	if err = bcrypt.CompareHashAndPassword(hp, []byte(cp)); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("cp - correct password")
	}

	wp := "wrong password"

	// Example of it not working
	if err = bcrypt.CompareHashAndPassword(hp, []byte(wp)); err != nil {
		fmt.Println(err)
		return
	}

}
