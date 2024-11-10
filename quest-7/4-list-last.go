package main

// Write a function ListLast that returns the Data interface of the last element of a linked list l.

func ListLast(l *List) interface{} {
	if l.Tail == nil {
		return nil
	}
	return l.Tail.Data
}
