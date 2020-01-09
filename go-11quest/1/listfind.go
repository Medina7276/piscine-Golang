package main

import (
	"fmt"
	
)

func main() {
	link := &List{}

	ListPushBack(link, "hello")
	ListPushBack(link, "hello1")
	ListPushBack(link, "hello2")
	ListPushBack(link, "hello3")

	found := ListFind(link, interface{}("hello2"), CompStr)

	fmt.Println(found)
	fmt.Println(*found)
}

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func CompStr(a, b interface{}) bool {
	return a == b
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

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	

	iter:=l.Head
	for iter != nil{
		if comp(iter.Data, ref){
			return &iter.Data
		}
		iter = iter.Next
	}
	return nil
}
