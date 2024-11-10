package main

// Write a function ListMerge that places elements of a list l2 at the end of another list l1.
// New elements should not be created

func ListMerge(l1 *List, l2 *List) {
	if l1.Head == nil {
		l1.Head = l2.Head
		l1.Tail = l2.Tail
	} else {
		l1.Tail.Next = l2.Head
		l1.Tail = l2.Tail
	}
}
