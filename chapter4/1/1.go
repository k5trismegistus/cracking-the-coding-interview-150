package main

import (
	"fmt"
)

type BiTree struct {
	Left  *BiTree
	Right *BiTree
}

func (t *BiTree) IsBalanced() bool {

	if t.Left == nil && t.Right == nil {
		return true
	}

	if t.Left == nil {
		return t.Right.Depth() <= 1
	}

	if t.Right == nil {
		return t.Left.Depth() <= 1
	}

	if !t.Left.IsBalanced() || !t.Right.IsBalanced() {
		return false
	}

	ld := t.Left.Depth()
	rd := t.Right.Depth()

	return ld-rd > -2 && ld-rd < 2
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

func main() {
	r := &BiTree{}
	r.Left = &BiTree{}
	r.Left.Left = &BiTree{}
	r.Right = &BiTree{}

	fmt.Println(r.IsBalanced())
}
