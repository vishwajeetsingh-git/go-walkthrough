package main

import "fmt" 

func main () {
	i := 2 
	if i % 2 == 0 {
		fmt.Println("even")
	}

	//Like `for` if statements can start with a short statement befor the condition.
	if j := 6; j % 2 == 0 { 
		fmt.Println("even")
	}
}
