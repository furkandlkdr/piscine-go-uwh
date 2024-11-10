package main

// Write a function ListClear that deletes all nodes from a linked list l.
// Tip: assign the list's pointer to nil

func ListClear(l *List) {
	current := l.Head
	for current != nil {
		next := current.Next
		current.Next = nil
		current = next
	}
	l.Head = nil
	l.Tail = nil
}
