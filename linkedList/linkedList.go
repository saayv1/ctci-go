package linkedList

import (
	"fmt"
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
	for n.Next != nil {
		*n = *n.Next
	}
}

func (n *Node) GoToTheBack() {
	for n.Previous != nil {
		*n = *n.Previous
	}
}

func (n *Node) Oscillate(){
	for n.Previous!=nil {
		fmt.Println(n.Previous," \t",n.Data,"\t",n.Next)
		*n= *n.Previous
	}
	for i:=0 ; i<=3; i++ {
                fmt.Println(n.Previous," \t",n.Data,"\t",n.Next)
                *n= *n.Next
        }
	fmt.Println(n.Previous," \t",n.Data,"\t",n.Next)

}

func Delete(n Node) {
	n1 := *n.Previous
	n2 := *n.Next
	n1.Next = &n2
	n2.Previous = &n1
}

