package main

import "fmt"
import "time"

func printHello() {
	for i := 0; i < 5; i++ {
		fmt.Printf("%d : Hello from Goroutine: %d\n",time.Now().UnixMilli(), i)
		time.Sleep(1000 * time.Millisecond)
	}
}

func printNumbers() {
	for i := 0; i < 5; i++ {
		fmt.Printf("%d : Goroutine: %d\n",time.Now().UnixMilli(), i)
		time.Sleep(1000 * time.Millisecond)
	}
}

func main() {
	fmt.Println("Goroutines Demo")

	go printHello()
	printNumbers()
}