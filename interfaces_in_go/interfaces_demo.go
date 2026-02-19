package main

import "fmt"

type Shape interface {
	Area() float64
	shapeName() string
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

func (c Circle) shapeName() string {
	return "Circle"
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) shapeName() string {
	return "Rectangle"
}

func printArea(s Shape) {
	fmt.Printf("Area of %s: %.2f\n", s.shapeName(), s.Area())
}

func main() {
	// Create instances of Circle and Rectangle

	circle := Circle{radius: 5}
	printArea(circle)

	rectangle := Rectangle{width: 4, height: 6}
	printArea(rectangle)
}
