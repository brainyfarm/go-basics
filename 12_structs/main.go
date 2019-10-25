package main

import (
	"fmt"
	"strconv"
)

// Person Define person struct
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

// Alien Short method struct
type Alien struct {
	homePlanet, superPower string
	age                    int
}

// Greeting method (value reciever)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer reciever)
func (p *Person) hasBirthday() {
	p.age++
}

func main() {
	// Init person using struct
	person1 := Person{
		firstName: "Olawale",
		lastName:  "Akinseye",
		city:      "Lagos",
		gender:    "m",
		age:       90,
	}

	person2 := Person{"Whitney", "Houston", "Texas", "f", 36}

	fmt.Println(person1)
	fmt.Println(person2)

	// Single field
	fmt.Println(person1.firstName)
	person1.hasBirthday()
	fmt.Println(person1.greet())

	fmt.Println(Alien{"Mercury", "FireMan", 22})
}
