package main

import (
	"fmt"
)

type Stack struct {
	Head *Node
}

type Node struct {
	Value int
	Next  *Node
}

func (s *Stack) Push(v int) {
	n := &Node{Value: v}

	n.Next = s.Head
	s.Head = n
}

func (s *Stack) Pop() int {
	n := s.Head
	s.Head = n.Next

	return n.Value
}

func (s *Stack) Peek() int {
	return s.Head.Value
}

func (s *Stack) IsEmpty() bool {
	return s.Head == nil
}

func (s *Stack) HasOnlyOne() bool {
	if s.IsEmpty() {
		return false
	}

	v := s.Pop()
	r := s.IsEmpty()
	s.Push(v)
	return r
}

func (s *Stack) Reverse() *Stack {
	ns := &Stack{}

	for {
		if s.Head == nil {
			break
		}

		v := s.Pop()
		ns.Push(v)
	}

	return ns
}

func (s *Stack) Sort() *Stack {
	if s.IsEmpty() || s.HasOnlyOne() {
		return s
	}

	left := &Stack{}
	right := &Stack{}
	pivot := s.Peek()

	n := s.Head
	for {
		if n == nil {
			break
		}

		if n.Value < pivot {
			left.Push(n.Value)
		} else {
			right.Push(n.Value)
		}

		n = n.Next
	}

	left = left.Sort().Reverse()
	right = right.Sort()

	ln := left.Head
	for {
		if ln == nil {
			break
		}

		right.Push(ln.Value)
		ln = ln.Next
	}

	return right
}

func (s *Stack) Print() {
	n := s.Head
	for {
		if n == nil {
			fmt.Println()
			break
		}

		fmt.Print(n.Value)
		n = n.Next
	}
}

func main() {
	s := &Stack{}
	s.Push(4)
	s.Push(5)
	s.Push(7)
	s.Push(2)
	s.Push(1)
	s.Push(3)
	s.Push(6)

	ss := s.Sort()
	ss.Print()
}
