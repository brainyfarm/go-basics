package main

import "fmt"

func main() {
	a := 5
	b := &a

	fmt.Println(a, b)
	fmt.Printf("a is %T b is %T\n", a, b)

	// Reading value from memory address
	fmt.Println("b is", *b)

	// Change value using pointer, b is a pointer to a ;)
	*b = 10
	fmt.Println("a is", a)
}
