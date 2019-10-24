package main

import "fmt"

func main() {
	x := 50
	y := 10

	if x < y {
		fmt.Printf("%d is less than %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}

	// else if
	color := "blue"

	if color == "red" {
		fmt.Println("color is red")
	} else if color == "blue" {
		fmt.Println("color is blue")
	} else {
		fmt.Println("color is NOT blue or red")
	}

	switch color {
	case "red":
		fmt.Println("paint it red")
	case "blue":
		fmt.Println("it is blue")
	default:
		fmt.Println("NOT red or blue")
	}
}
