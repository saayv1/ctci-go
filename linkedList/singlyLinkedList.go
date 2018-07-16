package linkedList
//Linked list package

import "fmt"

type SNode struct{
	Data int
	Next *SNode
}


func (n *SNode) AddRearSingle(number int){
	for n.Next != nil {
		n=n.Next;
	}
	nv := SNode{number,nil}
	n.Next = &nv;
}

func (n *SNode) PrintAllElements(){
	for n.Next != nil {
		fmt.Println(n.Data);
		n=n.Next;
	}
	fmt.Println(n.Data);
}
