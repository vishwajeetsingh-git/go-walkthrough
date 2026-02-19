package main

import "fmt"

func Area(i interface{}) {
	switch v := i.(type) {
	case Circle:
		fmt.Printf("Area of Circle: %.2f\n", 3.14*v.radius*v.radius)
	case Rectangle:
		fmt.Printf("Area of Rectangle: %.2f\n", v.width*v.height)
	default:
		fmt.Println("Unknown shape")
	}
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	width  float64
	height float64
}

func main() {
	circle := Circle{radius: 5}
	Area(circle)

	rectangle := Rectangle{width: 4, height: 6}
	Area(rectangle)
}
