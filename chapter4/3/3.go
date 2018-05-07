package main

import (
	"fmt"
)

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

func main() {
	sorted := []int{1, 2, 3, 4, 5, 6, 7}

	bitree := buildBiTree(sorted)

	fmt.Println(bitree.Depth())

}
