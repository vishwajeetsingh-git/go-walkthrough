
package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	current := 0
	next := 1
	return func() int {
		fib := current
		current = next
		next = fib + next
		return fib
	}
}

func fib(n int) map[int]int {
    current := 0
    next := 1
    sequence := make(map[int]int)
	fmt.Println("n:", n)
    for i:=0; i <= n; i++ {
		fmt.Printf("i: %d, current: %d, next: %d\n", i, current, next)
        sequence[i] = current
        current = next
		next = sequence[i] + next

    }

    

    return sequence

}


func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}


func main() {
	// f := fibonacci()
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(f())
	// }

	fmt.Println(fib(10))
}



