package main

type Stack struct {
	top  *Element
	size int
}

type Element struct {
	value interface{} // All types satisfy the empty interface, so we can store anything here.
	next  *Element
}

// Return the stack's length
func (s *Stack) Len() int {
	return s.size
}

// Push a new element onto the stack
func (s *Stack) Push(value interface{}) {
	s.top = &Element{value, s.top}
	s.size++
}

// Remove the top element from the stack and return it's value
// If the stack is empty, return nil
func (s *Stack) Pop() (value interface{}) {
	if s.size > 0 {
		value, s.top = s.top.value, s.top.next
		s.size--
		return
	}
	return nil
}

// // package main

// // type Node struct {
// // }
// // type List struct {
// // 	stack *Node
// // 	limit *Node
// // }

// // func (L *List) Push(data string) {
// // 	list := &Node{}
// // 	if L.stack < L.limit {

// // 	}

// // }

// package main

// import (
// 	"fmt"
// )

// type Node struct {
// 	Value string
// }

// func (n *Node) String() string {
// 	return fmt.Sprint(n.Value)
// }

// // NewStack returns a new stack.
// func NewStack() *Stack {
// 	return &Stack{}
// }

// // Stack is a basic LIFO stack that resizes as needed.
// type Stack struct {
// 	nodes []*Node
// 	count int
// }

// // Push adds a node to the stack.
// func (s *Stack) Push(n *Node) {
// 	s.nodes = append(s.nodes[:s.count], n)
// 	s.count++
// }

// // Pop removes and returns a node from the stack in last to first order.
// func (s *Stack) Pop() *Node {
// 	if s.count == 0 {
// 		return nil
// 	}
// 	s.count--
// 	return s.nodes[s.count]
// }

// // NewQueue returns a new queue with the given initial size.
// // func NewQueue(size int) *Queue {
// // 	return &Queue{
// // 		nodes: make([]*Node, size),
// // 		size:  size,
// // 	}
// // }
