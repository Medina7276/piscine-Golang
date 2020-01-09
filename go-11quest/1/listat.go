package main

import (
	"fmt"
	
)

func main() {
	link := &List{}

	ListPushBack(link, "hello")
	ListPushBack(link, "how are")
	ListPushBack(link, "you")
	ListPushBack(link, 1)

	fmt.Println(ListAt(link.Head, 6))
	fmt.Println(ListAt(link.Head, 3).Data)
	fmt.Println(ListAt(link.Head, 7))
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


func ListAt(l *NodeL, pos int) *NodeL{

	count := 1

	iter := l // просматриваем весь лист 
	for iter != nil { // проверяем пустой ли он
		if count == pos { //проверяем позицию
			return iter  //возвращаем данные
		}
		
		iter = iter.Next // сслыемся на след нод
		count++	//проходимся по всему листу 
	}
	return nil // если это конец листа верни мне пустоту
}
