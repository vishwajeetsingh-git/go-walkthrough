package main

import "fmt"

//Functions are values too

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}


//compute takes a function as an argument 
//calls the function with 3 and 4 as arguments
//returns the result of the function call
func compute(fn func(int, int) int) int {
	return fn(3, 4)
}

func main() {	
	fmt.Println(compute(add))
	fmt.Println(compute(sub))
}