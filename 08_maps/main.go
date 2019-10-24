package main

import "fmt"

func main() {
	// Define map
	emails := make(map[string]string)

	// Define and assign values
	ages := map[string]int{
		"Olawale": 12,
		"Sharon":  4,
	}

	fmt.Println(ages)

	// Assign key values
	emails["Olawale"] = "olawale@go.net"
	emails["Akinseye"] = "akinseye@go.net"
	emails["Trust"] = "trust@go.net"

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Olawale"])

	// Delete from a map
	delete(emails, "Trust")
	fmt.Println(len(emails))
	fmt.Println(emails)
}
