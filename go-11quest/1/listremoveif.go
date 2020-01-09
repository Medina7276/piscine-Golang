package main

import (
	"fmt"
)

func PrintList(l *List) {
	it := l.Head
	for it != nil {
		fmt.Print(it.Data, " -> ")
		it = it.Next
	}

	fmt.Print(nil, "\n")
}

func main() {
	link := &List{}
	link2 := &List{}

	fmt.Println("----normal state----")
	ListPushBack(link2, 1)
	PrintList(link2)
	ListRemoveIf(link2, 1)
	fmt.Println("------answer-----")
	PrintList(link2)
	fmt.Println()
	fmt.Println("----normal state----")
	ListPushBack(link, 1)
	ListPushBack(link, "Hello")
	ListPushBack(link, 1)
	ListPushBack(link, "There")
	ListPushBack(link, 1)
	ListPushBack(link, 1)
	ListPushBack(link, "How")
	ListPushBack(link, 1)
	ListPushBack(link, "are")
	ListPushBack(link, "you")
	ListPushBack(link, 1)
	PrintList(link)

	ListRemoveIf(link, 1)
	fmt.Println("------answer-----")
	PrintList(link)
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


func ListRemoveIf(l *List, data_ref interface{}) {
	l2 := &List{}
	iter := l.Head
	for iter != nil {
		if iter.Data == data_ref {
			iter = iter.Next
			continue
		}
		ListPushBack(l2, iter.Data)
		iter = iter.Next
	}
	*l = *l2
}
