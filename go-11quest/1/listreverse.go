package main

import (
	"fmt"
	
)

func main() {
	link := &List{}

	ListPushBack(link, 1)
	ListPushBack(link, 2)
	ListPushBack(link, 3)
	ListPushBack(link, 4)

	ListReverse(link)

	it := link.Head

	for it != nil {
		fmt.Println(it.Data)
		it = it.Next
	}

	fmt.Println("Tail", link.Tail)
	fmt.Println("Head", link.Head)
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

func ListPushFront(l *List, data interface{}) {

	n:= &NodeL{Data:data}

	if l.Head == nil {
		l.Head = n 
		return
	}

	n.Next = l.Head
	l.Head = n
}


func ListReverse(l *List) {

	n:= &List{} // создаем новый лист 

	iter := l.Head

	for iter != nil { //если след нод не равен 0 
		ListPushFront(n, iter.Data) // n= новый лист  iter.Data = ноды первого листа ( впихиваем в новый пустой лист n по одной ноде)
		iter = iter.Next  // иди к след ноду 
	}

	*l = *n // l =  это указатель и теперь он указывает на новый лист с перевернытуми нодами
}
