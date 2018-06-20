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

func PrintStuffL2R(np Node){
	for np.Previous != nil {
		np = *np.Previous
	}
	for np.Next !=nil {
		fmt.Println(np.Data)
		np = *np.Next
	}
}
