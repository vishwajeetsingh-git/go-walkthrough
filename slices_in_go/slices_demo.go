package main

import "fmt"

func main() {
	// Create a slice of int
	numbers := []int{1, 2, 3, 4, 5}

	// Print the slice
	fmt.Println("Numbers:", numbers)

	// slice from an array
	arr := [5]int{10, 20, 30, 40, 50}
	var sliceFromArr []int = arr[1:4] // This creates a slice that includes elements from index 1 to 3 (20, 30, 40)
	fmt.Println("Slice from array:", sliceFromArr)
}
