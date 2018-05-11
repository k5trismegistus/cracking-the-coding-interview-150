package main

import (
	"fmt"
)

type BiTree struct {
	Value int
	Left  *BiTree
	Right *BiTree
}

func (root *BiTree) Path(node *BiTree) []*BiTree {
	paths := [][]*BiTree{[]*BiTree{root}}

	for {
		newPaths := [][]*BiTree{}
		for _, p := range paths {
			if p[len(p)-1].Left != nil {
				nlp := append(p, p[len(p)-1].Left)

				if p[len(p)-1].Left == node {
					return nlp
				}

				newPaths = append(newPaths, nlp)
			}

			if p[len(p)-1].Right != nil {
				nrp := append(p, p[len(p)-1].Right)
				if p[len(p)-1].Right == node {
					return nrp
				}

				newPaths = append(newPaths, nrp)
			}

		}
		if len(newPaths) == 0 {
			return nil
		}
		paths = newPaths
	}
}

func (b *BiTree) CommonAncestor(n1, n2 *BiTree) *BiTree {
	path1 := b.Path(n1)
	path2 := b.Path(n2)

	fmt.Println(path1)
	fmt.Println(path2)

	for i := range path1 {
		if i >= len(path2) {
			if path1[i-1] == path2[i-1] {
				return path1[i-1]
			} else {
				return nil
			}
		}

		if path1[i] != path2[i] {
			return path1[i-1]
		}
	}
	return nil
}

func main() {
	root := &BiTree{Value: 5}
	root.Left = &BiTree{Value: 4}
	root.Right = &BiTree{Value: 8}
	root.Right.Left = &BiTree{Value: 6}

	ca := root.CommonAncestor(root.Right.Left, root.Right)
	fmt.Println(ca.Value)
}
