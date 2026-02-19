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

// we can declare methods on non-struct types as well
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	s := Square{side: 5}
	fmt.Printf("Area: %f\n", s.Area())
	fmt.Printf("Perimeter: %f\n", s.Perimeter())
	fmt.Printf("Diagonal: %f\n", s.Diagonal())

	f := MyFloat(-3.14)
	fmt.Printf("Absolute value: %f\n", f.Abs())
}