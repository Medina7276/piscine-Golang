package main

import (
	
	"fmt"
)

func PrintElem(node *NodeL) {
	fmt.Println(node.Data)
}

func StringToInt(node *NodeL) {
	node.Data = 2
}

func PrintList(l *List) {
	it := l.Head
	for it != nil {
		fmt.Print(it.Data, "->")
		it = it.Next
	}
	fmt.Print("nil","\n")
}

func main() {
	link := &List{}

	ListPushBack(link, 1)
	ListPushBack(link, "hello")
	ListPushBack(link, 3)
	ListPushBack(link, "there")
	ListPushBack(link, 23)
	ListPushBack(link, "!")
	ListPushBack(link, 54)

	PrintList(link)

	fmt.Println("--------function applied--------")
	ListForEachIf(link, PrintElem, IsPositive_node)

	ListForEachIf(link, StringToInt, IsNotNumeric_node)

	fmt.Println("--------function applied--------")
	PrintList(link)

	fmt.Println()
}

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func IsPositive_node(node *NodeL) bool {
	switch node.Data.(type) {
	case int, float32, float64, byte:
		return node.Data.(int) > 0
	case string, rune:
		return false
	}
	return false
}

func IsNegative_node(node *NodeL) bool {
	switch node.Data.(type) {
	case int, float32, float64, byte:
		return node.Data.(int) < 0
	case string, rune:
		return false
	}
	return false
}

func IsNotNumeric_node(node *NodeL) bool {
	switch node.Data.(type) {
	case int, float32, float64, byte:
		return false
	case string, rune:
		return true
	}
	return true
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

func ListForEachIf(l *List, f func(*NodeL), cond func(*NodeL) bool) {

	iter := l.Head 
	for iter != nil {
		if cond(iter) {
			f(iter)  // отправляеь хэд в функйию f и с помощью swith  он определяет какой тип интерфейса 
		}

		iter = iter.Next
	}

}
