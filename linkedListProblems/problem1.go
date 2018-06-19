package main

import (
	"../linkedList"
	"fmt"
)


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
	np := node
	for _,s:= range sli {
		np,_ =np.AddToFront(s)
		np.GoToTheFront()
	}
//	np.GoToTheBack()
	return *np
}

func main() {
	input :=[]int {1,2,3,5}
	nv:= addStuff(input)	
	nv.Oscillate()
	
	

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
