// The code snippet defines a basic shape interface and calculates the area for different shapes. However, it lacks a proper mechanism to handle different shape types using a type switch.
// Refactor the code to incorporate a type switch for handling special cases.
package main

import "fmt"

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width, Height float64
}

// TODO: Implement Area() method for Circle and Rectangle

func main() {
	shapes := []Shape{
		Circle{Radius: 3},
		Rectangle{Width: 4, Height: 5},
		// TODO: Add more shapes
	}

	// Print the area of each shape
	for _, shape := range shapes {
		// TODO: Use a type switch to handle different shape types
		fmt.Println("Area:", shape.Area())
	}
}
