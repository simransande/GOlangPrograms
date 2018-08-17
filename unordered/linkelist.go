package main

import "fmt"

type Node struct {
	prev *Node
	next *Node
	key  interface{}
}

type List struct {
	head *Node
	tail *Node
}

func (L *List) Insert(key interface{}) {
	list := &Node{
		next: L.head,
		key:  key,
	}
	if L.head != nil {
		L.head.prev = list
	}
	L.head = list

	l := L.head
	for l.next != nil {
		l = l.next
	}
	L.tail = l
}

func (l *List) Display() {
	list := l.head
	for list != nil {
		fmt.Printf("%v ", list.key)
		list = list.next
	}
	fmt.Println()
}

func (l *List) Search(data string) bool {
	list := l.head
	for list != nil {
		if list.key == data {
			return true
		}

		list = list.next

	}
	return false

}

func (l *List) Delete(data1 string) {
	list := l.head
	list1 := l.head
	if list.key == data1 {
		l.head = list.next
	} else {
		for list == list1 {
			list1 = list.next
		}
		if list.key == data1 {
			list1.next = list.next
		}
	}

}

// func Display(list *Node) {
// 	for list != nil {
// 		fmt.Printf("%v ->", list.key)
// 		list = list.next
// 	}
// 	fmt.Println()
// }
