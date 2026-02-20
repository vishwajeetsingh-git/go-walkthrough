package main

import "fmt"

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}


func (l *List[T]) Append(v T) *List[T] {
	if l == nil {
		return &List[T]{val: v}
	}

	current := l
	for current.next != nil {
		current = current.next
	}
	current.next = &List[T]{val: v}
	return l
}


func (l *List[T]) PrintList() {
    for current := l; current != nil; current = current.next {
        fmt.Println(current.val)
    }
}



func main() {
	fmt.Println("Generic List Demo")
	var l *List[int]
	l = l.Append(1)
	l = l.Append(2)
	l = l.Append(3)
	l.PrintList()
}