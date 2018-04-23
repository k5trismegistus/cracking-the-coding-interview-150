package main

import (
	"fmt"
)

type HanoiTower [3]*Stack

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

// Move top n discs on tower i to tower j
func (h *HanoiTower) Move(i, j, n int) {
	// rest tower
	k := 3 - (i + j)

	if n == 1 {
		d := h[i].Pop()
		h[j].Push(d)
		return
	}

	h.Move(i, k, n-1)
	h.Move(i, j, 1)
	h.Move(k, j, n-1)

	h.Print()
	fmt.Println("====================")
}

func (h *HanoiTower) Solve() {
	height := h[0].Len()

	h.Move(0, 2, height)
}

func (h *HanoiTower) Print() {
	for i := range h {
		n := h[i].Head
		for {
			if n == nil {
				fmt.Println()
				break
			}

			fmt.Print(n.Value)
			n = n.Next
		}
	}
}

func main() {
	h := &HanoiTower{&Stack{}, &Stack{}, &Stack{}}
	h[0].Push(9)
	h[0].Push(8)
	h[0].Push(7)
	h[0].Push(6)
	h[0].Push(5)
	h[0].Push(4)
	h[0].Push(3)
	h[0].Push(2)
	h[0].Push(1)

	h.Solve()
}
