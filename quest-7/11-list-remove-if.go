package main

// Write a function ListRemoveIf that removes all elements that are equal to the data_ref in the argument of the function.

func ListRemoveIf(l *List, data_ref interface{}) {
	current := l.Head
	var prev *NodeL
	for current != nil {
		if current.Data == data_ref {
			if current == l.Head {
				l.Head = current.Next
				current = l.Head
			} else {
				prev.Next = current.Next
				current = current.Next
			}
		} else {
			prev = current
			current = current.Next
		}
	}
	if l.Head == nil {
		l.Tail = nil
	}
}
