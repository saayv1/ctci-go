package main

import(
"fmt";
"../linkedList"
)

func main(){
	nv:= linkedList.InitializeNode(13)
	nc,err := nv.AddToBack(14)
	nc.GoToTheFront()
	fmt.Println(nc)
	fmt.Println(nv)	
	nv.GoToTheBack()
	fmt.Println(nv)
	if err != nil {
		fmt.Println(err)
	}

}
