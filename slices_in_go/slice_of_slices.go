package main

import "fmt"

func main() { 
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	board[0][0] = "X"	
	board[2][2] = "O"
	
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", board[i])
	}
}