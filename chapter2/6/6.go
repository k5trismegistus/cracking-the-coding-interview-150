package main

import (
	"fmt"
)

type LinkedList struct {
	head *LinkedListNode
}

type LinkedListNode struct {
	Value    int
	NextNode *LinkedListNode
}

func (ll *LinkedList) Add(nn *LinkedListNode) {
	n := ll.head
	if n == nil {
		ll.head = nn
		return
	}

	for {
		if n.NextNode == nil {
			n.NextNode = nn
			break
		}
		n = n.NextNode
	}
}

func createLinkedList(a []int) *LinkedList {
	ll := &LinkedList{}
	for _, v := range a {
		ll.Add(&LinkedListNode{Value: v})
	}

	return ll
}

func In(nl []*LinkedListNode, n *LinkedListNode) bool {
	if len(nl) == 0 {
		return false
	}

	for _, nnl := range nl {
		if nnl == n {
			return true
		}
	}
	return false
}

func (l *LinkedList) LastNode() *LinkedListNode {
	n := l.head

	if n == nil {
		return nil
	}

	for {
		if n.NextNode == nil {
			return n
		}

		n = n.NextNode
	}
}

func FindCircular(l *LinkedList) *LinkedListNode {
	al := []*LinkedListNode{}
	m := l.head
	for {
		if In(al, m) {
			return m
		}

		if m.NextNode == nil {
			break
		}

		al = append(al, m)
		m = m.NextNode
	}

	return nil
}

func main() {
	l := createLinkedList([]int{2, 8, 7, 5, 1})
	start := l.head.NextNode
	l.LastNode().NextNode = start

	n := FindCircular(l)

	fmt.Println(n.Value)
}
