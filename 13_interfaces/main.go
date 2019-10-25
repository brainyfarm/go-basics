package main

import (
	"fmt"
	"math"
)

// Shape Define interface
type Shape interface {
	area() float64
}

// Circle struct
type Circle struct {
	x, y, radius float64
}

// Rectangle struct
type Rectangle struct {
	width, height float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	circle := Circle{0, 0, 5}
	rectangle := Rectangle{5, 10}
	fmt.Printf("Circle Area: %f\n", getArea(circle))
	fmt.Printf("Rectangle Area: %f\n", getArea(rectangle))
	fmt.Println("Hello World")
}