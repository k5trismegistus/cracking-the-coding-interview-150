package main

import (
	"fmt"
)

type BiTree struct {
	Value int
	Left  *BiTree
	Right *BiTree
}

func (bt *BiTree) IsBiTree() (bool, int, int) {
	isBiTree := true
	var biggest int
	var smallest int
	var leftBiggest int
	var rightSmallest int

	if bt.Left != nil {
		isBiTree, leftBiggest, smallest = bt.Left.IsBiTree()

		if leftBiggest > bt.Value {
			isBiTree = false
		}
	} else {
		smallest = bt.Value
	}

	if bt.Right != nil {
		isBiTree, biggest, rightSmallest = bt.Right.IsBiTree()

		if rightSmallest < bt.Value {
			isBiTree = false
		}
	} else {
		biggest = bt.Value
	}

	return isBiTree, biggest, smallest
}

func main() {
	root := &BiTree{Value: 5}
	root.Left = &BiTree{Value: 4}
	root.Right = &BiTree{Value: 8}
	root.Right.Left = &BiTree{Value: 6}

	isBiTree, _, _ := root.IsBiTree()

	fmt.Println(isBiTree)
}
