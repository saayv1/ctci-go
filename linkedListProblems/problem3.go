package main

import (
	"../linkedList"
	"fmt"
)

func main() {
	n := linkedList.Node{nil, 0, nil}
	for i := 1; i < 100; i++ {
		n.AddFront(i)
	}
	k := 14
	nv := linkedList.GoBack(n)
	fmt.Println(nv.Data)
	fmt.Println("deleting the 14th node")

	for i := 0; i < k; i++ {
		if nv.Next == nil {
			fmt.Println("cannot delete the kth node")
			return
		}
		nv = nv.Next
	}
	nv = linkedList.Delete(nv)

	linkedList.PrintStuffL2R(*nv)

}
