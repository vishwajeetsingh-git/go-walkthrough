package main

import "fmt"
import "unsafe"

func main() {
	i := 1
	p := &i

	fmt.Println("i = ", i)
	fmt.Println("p = ", p)
	fmt.Println("*p = ", *p)

	*p = 2
	fmt.Println("i = ", i)

	var b *int

	b = &i
	fmt.Println("b = ", b)
	fmt.Println("*b = ", *b)
}