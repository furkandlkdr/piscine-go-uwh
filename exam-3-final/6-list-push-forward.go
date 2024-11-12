package main

// type NodeL struct {
// 	Data interface{}
// 	Next *NodeL
// }

// type List struct {
// 	Head *NodeL
// 	Tail *NodeL
// } Commented because of the conflict with the same type in the main package

func ListPushFront(l *List, data interface{}) {
	var myNode NodeL
	myNode.Data = data
	if l.Head == nil {
		l.Tail = &myNode
		l.Head = &myNode
		return
	}
	myNode.Next = l.Head
	l.Head = &myNode
}
