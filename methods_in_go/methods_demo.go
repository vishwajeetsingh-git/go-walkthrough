package main

import "fmt"

type Square struct {
	side float64
}

func (s Square) Area() float64 {
	return s.side * s.side
}

func (s Square) Perimeter() float64 {
	return 4 * s.side
}

func (s Square) Diagonal() float64 {
	return s.side * 1.41421356237
}

func main() {
	s := Square{side: 5}
	fmt.Printf("Area: %f\n", s.Area())
	fmt.Printf("Perimeter: %f\n", s.Perimeter())
	fmt.Printf("Diagonal: %f\n", s.Diagonal())
}