package linkedList

type node struct {
	next *node
	data int
}

func InitializeNode(value int)node {
	nv := node{nil,value};
	return nv
}	


