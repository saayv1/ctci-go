package linkedList
import(
"errors"
)
type node struct {
	Next *node
	Previous *node
	Data int
}

func InitializeNode(value int) node {
	nv := node{nil,nil, value}
	return nv
}

func (n *node)AddToBack( value int)(node,error){
	var nv node
	if n.Previous != nil {
		return nv, errors.New("Cannot append")
	}
	nv = node{n,nil,value}
	n.Previous = &nv
	return nv,nil
}

func(n *node)AddToFront(value int)(node,error){
	var nv node
	if n.Next!=nil	{
		return nv,errors.New("Cannot append to the rear")
	}
	nv = node{nil,n,value}
	n.Next = &nv
	return nv,nil
}

func (n *node)GoToTheFront(){
	for {
		if n.Next == nil {
			return;
		} else {
			*n=*n.Next;
		}
	}
}


func (n *node)GoToTheBack(){
	for{
		if n.Previous == nil {
			return
		} else {
			*n = *n.Previous
		}
	}
}
