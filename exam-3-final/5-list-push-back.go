package main

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	var myNode NodeL
	myNode.Data = data
	if l.Head == nil {
		l.Head = &myNode
		l.Tail = &myNode
		return
	}
	current := l.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = &myNode
	l.Tail = &myNode

}
