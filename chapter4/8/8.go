package main

import (
	"fmt"
)

type BiTree struct {
	Value int
	Left  *BiTree
	Right *BiTree
}

func (t *BiTree) ToString() string {
	str := fmt.Sprintf("%d", t.Value)

	if t.Left == nil {
		str += "n"
	} else {
		str += t.Left.ToString()
	}

	if t.Right == nil {
		str += "n"
	} else {
		str += t.Right.ToString()
	}

	return str
}

func (t *BiTree) Match(s *BiTree) bool {
	return t.ToString() == s.ToString()
}

func (t *BiTree) SubTree(s *BiTree) bool {
	if t.Value == s.Value && t.Match(s) {
		return true
	}

	if t.Left != nil && t.Left.SubTree(s) {
		return true
	}

	if t.Right != nil && t.Right.SubTree(s) {
		return true
	}

	return false
}

func main() {
	root := &BiTree{Value: 5}
	root.Left = &BiTree{Value: 4}
	root.Right = &BiTree{Value: 8}
	root.Right.Left = &BiTree{Value: 6}

	sub := &BiTree{Value: 8}
	// sub.Left = &BiTree{Value: 6}

	fmt.Println(root.SubTree(sub))
}
