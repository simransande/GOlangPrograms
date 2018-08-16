package main

import "fmt"

// type Node struct {
// 	prev *Node
// 	next *Node
// 	currentNode.Value=0
// }
// type SinglyLinkedList struct {
// 	front *Node

// 	length int
// }

// func (s *SinglyLinkedList) Prepend(n *Node) {
// 	if s.front == nil {
// 		s.front = n
// 	} else {
// 		n.next = s.front
// 		s.front = n
// 	}

// 	s.length++
// }

// func (s *SinglyLinkedList) Remove(n *Node) {
// 	if s.front == n {
// 		s.front = n.next
// 		s.length--
// 	} else {
// 		currentNode := s.front

// 		// search for node n
// 		for currentNode != nil && currentNode.next != nil && currentNode.next != n {
// 			currentNode = currentNode.next
// 		}

// 		// see if current's next node is n
// 		// if it's not n, then node n wasn't found in list s
// 		if currentNode.next == n {
// 			currentNode.next = currentNode.next.next
// 			s.length--
// 		}
// 	}
// }

// func (s *SinglyLinkedList) Find(value interface{}) *Node {
// 	currentNode := s.front
// 	for currentNode != nil && currentNode.Value != value && currentNode.next != nil {
// 		currentNode = currentNode.next
// 	}

// 	return currentNode
// }

// // Length returns the amount of nodes in list s
// func (s *SinglyLinkedList) Length() int {
// 	return s.length
// }

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
		fmt.Printf("%+v ->", list.key)
		list = list.next
	}
	fmt.Println()
}

func search(data) {
	list := &Node{}
}

func Display(list *Node) {
	for list != nil {
		fmt.Printf("%v ->", list.key)
		list = list.next
	}
	fmt.Println()
}

func ShowBackwards(list *Node) {
	for list != nil {
		fmt.Printf("%v <-", list.key)
		list = list.prev
	}
	fmt.Println()
}

func (l *List) Reverse() {
	curr := l.head
	var prev *Node
	l.tail = l.head

	for curr != nil {
		next := curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	l.head = prev
	Display(l.head)
}
