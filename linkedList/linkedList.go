package linkedList

import(
	"fmt"
)

type Node struct {
	Previous *Node
	Data int
	Next *Node
}

func (n *Node) AddRear(value int){
	for n.Previous != nil {
		n = n.Previous
	}
	nv := Node{nil,value, nil}
	nv.Next = n;
	n.Previous = &nv
}

func (n *Node) AddFront(value int){
	for n.Next != nil {
		n = n.Next
	}
	nv:= Node{nil,value,nil}
	nv.Previous = n;
	n.Next = &nv
}

func AddStuffRear(n *Node,sli []int){
        for _,s := range sli {
                n.AddRear(s)
        }
}

func AddStuffFront(n *Node,sli []int){
        for _,s := range sli {
                n.AddFront(s)
        }
}

func Delete(n *Node)*Node{
	if n.Previous == nil && n.Next !=nil {
		np := n.Next
		np.Previous = nil
		n.Next=nil
		n.Previous = nil
		return np
	}else if n.Next == nil && n.Previous!=nil {
		np := n.Previous
		np.Next = nil
		n.Next = nil
		n.Previous = nil
		return np
	}else if n.Next != nil && n.Previous!=nil {
		n1 := n.Previous
		n2 := n.Next
		n1.Next = n2
		n2.Previous = n1
		n.Previous = nil
		n.Next = nil
		return n1
	}
	return n
}
		

func PrintStuffL2R(np Node){
	for np.Previous != nil {
		np = *np.Previous
	}
	for np.Next !=nil {
		fmt.Println(np.Data)
		np = *np.Next
	}
	fmt.Println(np.Data)
}

func GoBack(n Node)*Node{
	np := &n
	for np.Previous!=nil {
		np = np.Previous
	}
	return np
}
