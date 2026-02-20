package main

import (
	"fmt"
	"time"
)

// recives a channel and a number n 
// computes the first n numbers in the fibonacci 
// sequence and sends them to the channel 
func fibonacci(n int, ch chan int) {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		ch <- a
		a, b = b, a+b
	}
	close(ch)
}

func main() {
	fmt.Println("Goroutines Channels Demo")

	ch := make(chan string)

	go func() {
		time.Sleep(500 * time.Millisecond)
		ch <- "Hello from Goroutine"
	}()

	fmt.Println("Waiting for message from Goroutine...")
	msg := <-ch
	fmt.Println("Received message from Goroutine:")
	fmt.Println(msg)


	// channels can be buffered 
	bufferedCh := make(chan string, 2)
	bufferedCh <- "Message 1"
	bufferedCh <- "Message 2"

	fmt.Println("Messages in buffered channel:")
	fmt.Println(<-bufferedCh)
	fmt.Println(<-bufferedCh)

	// channels can be closed by the sender 
	// to indicate that no more vaues will be 
	// sent 

	ch2 := make(chan int)
	go fibonacci(10, ch2)

	fmt.Println("Fibonacci sequence:")

	// range over the channel to receive values
	//  until it is closed
	for num := range ch2 {
		fmt.Println(num)
	}
}