package main

import (
	"../linkedList"
	"fmt"
)

func printStuff(n linkedList.Node) {
	for {
		if n.Previous == nil {
			fmt.Println(n.Data)
			return
		} else {
			fmt.Println(n.Data)
			n = *n.Previous
		}
	}
}

func ps2(n linkedList.Node) {
	for n.Next!=nil {
		fmt.Println(n.Data)
		n=*n.Next
	}
	fmt.Println(n.Data)
	return
}

func addStuff(sli []int)linkedList.Node{
	node:= linkedList.InitializeNode(sli[0])
	for _,s:= range sli {
		node,_ = node.AddToFront(s)
		node.GoToTheFront()
	}
	//node.GoToTheBack();
	return *node
}

func main() {
	input :=[]int {2,4,2,5,6,7,8,6,5,6,3,1}
	nv:= addStuff(input)	
	nv.GoToTheBack()
	ps2(nv)
	

/*
	fmt.Println(nv.Data);
	nv= *nv.Previous
	fmt.Println(nv.Data);	

nv= *nv.Previous
fmt.Println(nv.Data);
nv= *nv.Previous
fmt.Println(nv.Data);
nv= *nv.Previous
fmt.Println(nv.Data);
*/
}
