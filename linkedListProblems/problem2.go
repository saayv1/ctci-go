package main

import (
	"../linkedList"
	"fmt"
)

func main() {
	k := 1
	n := linkedList.SNode{2, nil}
	n.AddRearSingle(12)
	n.AddRearSingle(22)
	n.AddRearSingle(32)
	n.PrintAllElements()

	p1 := &n
	p2 := &n

	for i := 0; i < k; i++ {
		if p1 == nil {
			fmt.Println("There is no k th to last element")
			return
		} else {
			p1 = p1.Next
		}
	}

	for p1.Next != nil {
		p1 = p1.Next
		p2 = p2.Next
	}
	fmt.Printf("\n\nThe answer is %d\n\n", p2.Data)

}
