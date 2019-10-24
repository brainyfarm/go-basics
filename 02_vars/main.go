package main

import (
	"fmt"
)

var name = "Olawale"

func main() {
	var age int32 = 37

	// Constant with implicit typing
	const isCool = true

	// Shorthand assignment, implicit typing
	country := "Nigeria"

	size := 1.3

	email, address := "brainbrain@gmail.com", "Somewhere in Nigeria"

	fmt.Println(name, age, isCool, country, size, email, address)
	fmt.Printf("age is an %T\n", age)
}
