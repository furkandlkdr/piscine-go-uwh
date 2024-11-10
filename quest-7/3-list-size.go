package main

// Write a function ListSize that returns the number of elements in a linked list l.

func ListSize(l *List) int {
	size := 0
	current := l.Head
	for current != nil {
		size++
		current = current.Next
	}
	return size
}
