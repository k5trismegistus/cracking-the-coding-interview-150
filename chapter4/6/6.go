package main

import (
	"fmt"
)

type BiTree struct {
	Value  int
	Parent *BiTree
	Left   *BiTree
	Right  *BiTree
}

func (b *BiTree) Next() *BiTree {
	if b.Right != nil {
		n := b.Right
		for {
			if n.Left == nil {
				return n
			}
			n = n.Left
		}
	} else {
		n := b
		for {
			if n.Parent == nil {
				return nil
			}

			if n.Parent.Left == n {
				return n.Parent
			}

			n = n.Parent
		}
	}
}

func main() {
	root := &BiTree{Value: 5}
	root.Left = &BiTree{Value: 4, Parent: root}
	root.Right = &BiTree{Value: 8, Parent: root}
	root.Right.Left = &BiTree{Value: 6, Parent: root.Right}

	fmt.Println(root.Right.Next().Value)
}
