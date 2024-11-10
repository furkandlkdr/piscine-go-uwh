package main

func IsPositiveNode(node *NodeL) bool {
	switch node.Data.(type) {
	case int, float32, float64, byte:
		return node.Data.(int) > 0
	default:
		return false
	}
}

func IsAlNode(node *NodeL) bool {
	switch node.Data.(type) {
	case int, float32, float64, byte:
		return false
	default:
		return true
	}
}

// Write a function ListForEachIf that applies a function given as argument to the data within some of the nodes of the list l.
// This function receives two functions:
// f is a function that is applied to the node.
// cond is a function that returns a boolean and it will be used to determine if the function f should be applied to the node.
// The function given as argument must have a pointer *NodeL as argument.

func ListForEachIf(l *List, f func(*NodeL), cond func(*NodeL) bool) {
	current := l.Head
	for current != nil {
		if cond(current) {
			f(current)
		}
		current = current.Next
	}
}
