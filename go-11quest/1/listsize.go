package main

import (
	"fmt"

)

func main() {
	link := &List{}

	ListPushFront(link, "Hello")
	ListPushFront(link, "2")
	ListPushFront(link, "you")
	ListPushFront(link, "man")
	

	fmt.Println(ListSize(link))
}

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}


func ListPushFront(l *List, data interface{}) {

	n:= &NodeL{Data:data}

	if l.Head == nil {
		l.Head = n 
		return
	}

	n.Next = l.Head
	l.Head = n
}


func ListSize(l *List) int {

	count := 0
	
	iter := l.Head
	
	for iter != nil {
		iter = iter.Next
	    count++	
	}
	return count

}

