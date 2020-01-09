package main

import (
	"fmt"
)

func main() {
	link := &List{}
	link2 := &List{}

	ListPushBack(link, "three")
	ListPushBack(link, 3)
	ListPushBack(link, "1")

	fmt.Println(ListLast(link))
	fmt.Println(ListLast(link2))
}


type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	
	n:= &NodeL{Data:data}
	
	if l.Head == nil {
		l.Head = n
		return
	}

	iter := l.Head
	for iter.Next != nil {
		iter = iter.Next
	}
	iter.Next = n
}


func ListLast(l *List) interface{} {
	
	if l.Head == nil {
		
		return nil
	}
	
	
	iter := l.Head
	for iter.Next != nil {  
		iter = iter.Next
	}
	//l.Tail = iter
	return iter.Data
}

 //iter = "three"
//iter.Next = 3
// iter = iter.Next
// 
// iter = 3 
// iter.Next = 1
// iter = iter.Next

// iter = 1
// iter.Next = nil

