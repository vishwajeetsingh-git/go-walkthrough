package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

type Employee struct {
	Name string
	Age  int
}

func main() {
	fmt.Println(Vertex{1, 2})

	//struct fields are accessed using a dot.
	e := Employee{"Vishwajeet", 25}
	fmt.Println(e.Name)
	fmt.Println(e.Age)

	//change struct fields
	e.Age = 26
	fmt.Println(e.Age)

	// structs fields can be accessed through a struct pointer.
	// the struct pointer is automatically dereferenced.
	p := &e
	fmt.Println("Name using Pointer to Struct e :",p.Name)
	fmt.Println("Age using Pointer to Struct e :",p.Age)
}
