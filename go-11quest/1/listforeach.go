package main

import (
	"fmt"

)

func main() {
	link := &List{}

	ListPushBack(link, "1")
	ListPushBack(link, 2)
	ListPushBack(link, "3")
	ListPushBack(link, "5")

	ListForEach(link, Subtract3_node)

	it := link.Head
	for it != nil {
		fmt.Println(it.Data)
		it = it.Next
	}
}

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}


func Add2_node(node *NodeL) {
	switch node.Data.(type) {
	case int:
		node.Data = node.Data.(int) + 2
	case string:
		node.Data = node.Data.(string) + "2"
	}
}

func Subtract3_node(node *NodeL) {
	switch node.Data.(type) {
	case int:
		node.Data = node.Data.(int) - 3
	case string:
		node.Data = node.Data.(string) + "3"
	}
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

func ListForEach(l *List, f func(*NodeL)) {

	iter := l.Head 
	for iter != nil {
		f(iter)  // отправляеь хэд в функйию f и с помощью swith  он определяет какой тип интерфейса 
		iter = iter.Next
	}
}

