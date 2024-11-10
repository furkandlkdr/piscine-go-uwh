package main

type NodeI struct {
	Data int
	Next *NodeI
}

//Write a function ListSort that sorts the nodes of a linked list by ascending order.

func ListSort(l *NodeI) *NodeI {
	if l == nil {
		return nil
	}
	var current *NodeI
	var next *NodeI
	for current = l; current != nil; current = current.Next {
		for next = current.Next; next != nil; next = next.Next {
			if current.Data > next.Data {
				current.Data, next.Data = next.Data, current.Data
			}
		}
	}
	return l
}
