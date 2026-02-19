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

	// slices are like references to arrays. When you create a slice from an array,
	// it does not copy the elements of the array. Instead,
	// it creates a reference to the original array.
	// This means that if you modify the elements of the slice,
	// it will also modify the corresponding elements in the original array.
	fmt.Println("Original array before modification:", arr)
	sliceFromArr[0] = 25 // This modifies the first element of the slice, which is also the second element of the original array
	fmt.Println("Modified slice from array:", sliceFromArr)
	fmt.Println("Original array after modification:", arr)

	//slice literals

	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)

}
