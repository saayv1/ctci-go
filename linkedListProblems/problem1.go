package main

import (
	"fmt"
	"../linkedList"
)

func DeleteDuplicates(n linkedList.Node)*linkedList.Node{
	np := &n
	m := make(map[int]int)
	np = linkedList.GoBack(*np)
	for np.Next != nil {
		x,ok := m[np.Data]
		if ok {
			x++;
			m[np.Data] = x
		}	else	{
			m[np.Data] = 1
		}
		np = np.Next
	}
	   x,ok := m[np.Data]
                if ok {
                        x++;
                        m[np.Data] = x
                }       else    {
                        m[np.Data] = 1
                }
	for key,value := range m {
		if value==1 {
			delete(m,key)
		}
	}
	fmt.Println(m)
	for key,_ := range m{
		np = linkedList.GoBack(*np)
		for np.Next != nil	{
			if np.Data == key {
				if m[key]>1 {
					m[key] -= 1
					np = linkedList.Delete(np)
				}
			}
			np = np.Next
		}
	}
	np = linkedList.GoBack(*np)
	return np
}


func main(){
	n := linkedList.Node{nil,1,nil}
	np := &n
	slice := []int{55,54,55,52,51,55}	
	linkedList.AddStuffFront(&n,slice)
	n.AddFront(15)
	n.AddRear(15)
	n.AddRear(17)
	n.AddFront(15)
	n.AddRear(17)
	np = linkedList.GoBack(*np)
	linkedList.PrintStuffL2R(*np)
	fmt.Println("\n\n\n")	
	np = DeleteDuplicates(*np)
	np = linkedList.GoBack(*np)
	linkedList.PrintStuffL2R(*np)
}
