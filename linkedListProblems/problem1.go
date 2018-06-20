package main

import (
	"../linkedList"
//	"fmt"
)

func main(){
	n := linkedList.Node{nil,1,nil}
	slice := []int{55,54,53,52,51,50}	
	linkedList.AddStuffFront(&n,slice)
	n.AddRear(14)
	n.AddRear(15)
	n.AddRear(17)
	linkedList.PrintStuffL2R(n)
}
