package main

import "fmt"

func main()	 {
	// for loop with a single condition
	for i := 0; i < 5; i++ {
		fmt.Printf("Iteration %d\n", i)
	}

	// for loop with no pre and post statements 
	j := 5
	for ; j > 0 ; {
		fmt.Printf("Countdown: %d\n", j)
		j--
	}
}