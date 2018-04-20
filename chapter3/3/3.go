package main

import (
	"fmt"
)

const stackMax = 3

type SetOfStacks []Stack

type Stack struct {
	Head *Node
}

type Node struct {
	Value int
	Next  *Node
}

func (s *Stack) Len() int {
	l := 0
	n := s.Head

	for {
		if n == nil {
			return l
		}

		l++
		n = n.Next
	}
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

func (ss *SetOfStacks) Push(v int) {
	if len(*ss) == 0 || (*ss)[0].Len() >= stackMax {
		fmt.Println(v)
		ns := &Stack{}

		ns.Push(v)
		*ss = append(SetOfStacks{*ns}, *ss...)

		return
	}

	(*ss)[0].Push(v)
}

func (ss *SetOfStacks) Pop() int {
	if len(*ss) == 0 {
		panic("SetOfStacks is empty")
	}

	s := (*ss)[0]

	r := s.Pop()

	if s.Len() == 0 {
		*ss = (*ss)[1:len(*ss)]
	}
	return r
}

func (ss *SetOfStacks) PopAt(i int) int {
	if len(*ss) < i {
		panic("Out of range")
	}

	return (*ss)[i].Pop()
}

func main() {
	ss := &SetOfStacks{}
	ss.Push(1)
	ss.Push(2)
	ss.Push(3)
	ss.Push(4)
	ss.Push(5)
	ss.Push(6)
	ss.Push(7)
	fmt.Println(ss)

	fmt.Println(ss.PopAt(1))
}
