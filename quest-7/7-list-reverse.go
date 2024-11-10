package main

// Write a function ListReverse that reverses the order of the elements of a given linked list l.

func ListReverse(l *List) {
	var prev *NodeL
	current := l.Head
	l.Tail = current
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	l.Head = prev
}
