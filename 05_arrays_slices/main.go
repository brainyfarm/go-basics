package main

import "fmt"

func main() {
	// Arrays explicit fix length, assign values later
	var fruitArr [2]string

	// Declare and assign
	carArr := [2]string{"Toyota", "Benz"}

	// Unknown length, array slices to the rescue
	carSlice := []string{"Toyota", "Benz", "BMW"}

	// Assign values
	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"

	fmt.Println(fruitArr)
	fmt.Println(fruitArr[1])
	fmt.Println(carArr)
	fmt.Println(carSlice)
	fmt.Println(len(carSlice))
	fmt.Println(carSlice[0:2])

}
