package main

import "fmt"

func main() {
	ids := []int{
		1,
		2,
		3,
	}

	networth := map[string]int{
		"Olawale": 12000,
		"Sharon":  400,
	}

	// Loop through id
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// Loop through id with using index (i)
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum", sum)

	// Range with map key and value
	for k, v := range networth {
		fmt.Printf("%s: %d\n", k, v)
	}

	// Range with map key only
	for name := range networth {
		fmt.Println("Name: " + name)
	}

}
