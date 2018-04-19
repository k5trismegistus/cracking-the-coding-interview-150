package main

import (
	"fmt"
)

type StackMin struct {
	head    *Node
	minHead *Node
}

type Node struct {
	Value int
	Next  *Node
}

func (s *StackMin) Push(v int) {
	n := &Node{Value: v}
	m := s.minHead

	n.Next = s.head
	s.head = n.Next

	if m == nil || v < m.Value {
		n.Next = s.minHead
		s.minHead = n
	} else {
		nn := &Node{Value: m.Value}
		nn.Next = s.minHead
		s.minHead = nn
	}
}

func (s *StackMin) Pop() int {
	r := s.head.Value

	s.head = s.head.Next
	s.minHead = s.minHead.Next

	return r
}

func (s *StackMin) Min() int {
	return s.minHead.Value
}

func main() {
	sm := &StackMin{}

	sm.Push(3)
	sm.Push(5)
	sm.Push(1)

	fmt.Println(sm.Min())
}
