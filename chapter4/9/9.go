package main

import (
	"fmt"
)

type BiTree struct {
	Value int
	Left  *BiTree
	Right *BiTree
}

type Path []*BiTree

func (p *Path) Sum() int {
	s := 0
	for _, b := range []*BiTree(*p) {
		s += b.Value
	}
	return s
}

func (p *Path) ToString() string {
	s := "["
	for i, b := range []*BiTree(*p) {
		if i == len([]*BiTree(*p))-1 {
			s += fmt.Sprintf("%d]", b.Value)
			break
		}
		s += fmt.Sprintf("%d, ", b.Value)
	}

	return s
}

func (b *BiTree) FindPath(v int) []Path {
	initPath := Path([]*BiTree{b})
	paths := []Path{initPath}
	goodPaths := []Path{}

	for {
		newPaths := []Path{}
		for _, p := range paths {
			if p.Sum() == v {
				goodPaths = append(goodPaths, p)
			}

			if p[len(p)-1].Left != nil && p.Sum() < v {
				nlp := append(p, p[len(p)-1].Left)

				newPaths = append(newPaths, nlp)
			}

			if p[len(p)-1].Right != nil && p.Sum() < v {
				nrp := append(p, p[len(p)-1].Right)

				newPaths = append(newPaths, nrp)
			}

		}
		if len(newPaths) == 0 {
			break
		}
		paths = newPaths
	}

	return goodPaths
}

func (b *BiTree) FindAllPath(v int) []Path {
	paths := b.FindPath(v)

	if b.Left != nil {
		paths = append(paths, b.Left.FindAllPath(v)...)
	}

	if b.Right != nil {
		paths = append(paths, b.Right.FindAllPath(v)...)
	}

	return paths
}

func main() {
	root := &BiTree{Value: 5}
	root.Left = &BiTree{Value: 4}
	root.Right = &BiTree{Value: 8}
	root.Right.Left = &BiTree{Value: 6}
	root.Left.Right = &BiTree{Value: 2}
	root.Left.Right.Left = &BiTree{Value: 3}
	root.Left.Right.Right = &BiTree{Value: 7}
	root.Left.Right.Right.Left = &BiTree{Value: 2}

	v := 9

	paths := root.FindAllPath(v)

	for _, p := range paths {
		fmt.Println(p.ToString())
	}
}
