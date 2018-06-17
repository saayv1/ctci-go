package linkedList

import (
	"errors"
)

type Node struct {
	Next     *Node
	Previous *Node
	Data     int
}

func InitializeNode(value int) *Node {
	nv := Node{nil, nil, value}
	return &nv
}

func (n *Node) AddToBack(value int) (*Node, error) {
	var nv Node
	if n.Previous != nil {
		return &nv, errors.New("Cannot append")
	}
	nv = Node{n, nil, value}
	n.Previous = &nv
	return &nv, nil
}

func (n *Node) AddToFront(value int) (*Node, error) {
	var nv Node
	if n.Next != nil {
		return &nv, errors.New("Cannot append to the rear")
	}
	nv = Node{nil, n, value}
	n.Next = &nv
	return &nv, nil
}

func (n *Node) GoToTheFront() {
	for {
		if n.Next == nil {
			return
		} else {
			*n = *n.Next
		}
	}
}

func (n *Node) GoToTheBack() {
	for {
		if n.Previous == nil {
			return
		} else {
			*n =*n.Previous
		}
	}
}

func Delete(n Node) {
	n1 := *n.Previous
	n2 := *n.Next
	n1.Next = &n2
	n2.Previous = &n1
}
