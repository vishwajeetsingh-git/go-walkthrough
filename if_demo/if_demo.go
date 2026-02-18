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

	//variable declared in an if statement are also accessible in any else if or else blocks that follow it.
	if k := 7; k % 2 == 0 {
		fmt.Println("even : ", k)
	} else {
		fmt.Println("odd : ", k)
	}

}
