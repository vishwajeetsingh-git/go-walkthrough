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

	// for is go's while
	k := 0 
	for k < 5 {
		fmt.Printf("While loop iteration %d\n", k)
		k++
	}

	//infinite loop
	// for {
	// 	fmt.Println("This will run forever")
	// }


	//range form of for loop to iterate over a slice
	pow := []int{2, 4, 8, 16, 32}
	for index, value := range pow {
		fmt.Printf("2^%d = %d\n", index, value)
	}
}