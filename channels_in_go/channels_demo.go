package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Goroutines Channels Demo")

	ch := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- "Hello from Goroutine"
	}()

	fmt.Println("Waiting for message from Goroutine...")
	msg := <-ch
	fmt.Println("Received message from Goroutine:")
	fmt.Println(msg)
}