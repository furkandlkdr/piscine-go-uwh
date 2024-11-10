package main

func CompStr(a, b interface{}) bool {
	return a == b
}

// Write a function ListFind that returns the address of the data interface of the first node in the list l that is determined to be equal to ref by the function CompStr.
// For this exercise the function CompStr must be used

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	current := l.Head
	for current != nil {
		if comp(current.Data, ref) {
			return &current.Data
		}
		current = current.Next
	}
	return nil
}
