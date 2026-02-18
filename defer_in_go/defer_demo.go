package main

import "fmt"

func main() {

	defer fmt.Println("Counting done")

	for i := 1; i <= 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("Counting...")
}	