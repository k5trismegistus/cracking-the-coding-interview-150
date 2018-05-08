package main

import (
	"fmt"
)

type LinkedList struct {
	Root *LinkedListNode
}

type LinkedListNode struct {
	Value int
	Next  *LinkedListNode
}

func (ll *LinkedList) Print() {
	n := ll.Root

	for {
		if n == nil {
			fmt.Print("\n")
			return
		}

		fmt.Print(n.Value)
		n = n.Next
	}
}

func (ll *LinkedList) Push(v int) {
	n := &LinkedListNode{Value: v}

	if ll.Root == nil {
		ll.Root = n
	} else {
		cn := ll.Root

		for {
			if cn.Next == nil {
				cn.Next = n
				return
			}
			cn = cn.Next
		}
	}
}

type BiTree struct {
	Value int
	Left  *BiTree
	Right *BiTree
}

func (t *BiTree) Depth() int {
	if t.Left == nil && t.Right == nil {
		return 1
	}

	var ld int
	var rd int

	if t.Left == nil {
		ld = 0
	} else {
		ld = t.Left.Depth()
	}

	if t.Right == nil {
		rd = 0
	} else {
		rd = t.Right.Depth()
	}

	if ld < rd {
		return rd + 1
	}
	return ld + 1
}

func buildBiTree(s []int) *BiTree {
	if len(s) == 1 {
		return &BiTree{Value: s[0]}
	}

	c := len(s) / 2
	root := s[c]
	rootNode := &BiTree{Value: root}

	rootNode.Left = buildBiTree(s[0:c])

	if len(s) > 2 {
		rootNode.Right = buildBiTree(s[c+1 : len(s)])
	}

	return rootNode
}

func (b *BiTree) buildSameDepthList() []*LinkedList {
	lists := []*LinkedList{}
	ns := []*BiTree{b}

	for {
		if len(ns) == 0 {
			break
		}

		nextNs := []*BiTree{}
		ll := &LinkedList{}
		for _, n := range ns {
			ll.Push(n.Value)

			if n.Left != nil {
				nextNs = append(nextNs, n.Left)
			}

			if n.Right != nil {
				nextNs = append(nextNs, n.Right)
			}
		}
		lists = append(lists, ll)
		ns = nextNs
	}

	return lists
}

func main() {
	sorted := []int{1, 2, 3, 4, 5, 6, 7}

	bitree := buildBiTree(sorted)

	lists := bitree.buildSameDepthList()

	for _, l := range lists {
		l.Print()
	}
}
