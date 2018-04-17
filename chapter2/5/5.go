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

func Tashizan(l1, l2 *LinkedList) int {
	r := 0
	m := 1

	n1 := l1.head
	n2 := l2.head

	for {

		v1 := 0
		v2 := 0

		if n1 != nil {
			v1 = n1.Value
		} else {
			v1 = 0
		}

		if n2 != nil {
			v2 = n2.Value
		} else {
			v2 = 0
		}

		r = r + (v1+v2)*m

		if n1 != nil {
			n1 = n1.NextNode
		}
		if n2 != nil {
			n2 = n2.NextNode
		}

		if n1 == nil && n2 == nil {
			break
		}

		m = m * 10
	}

	return r
}

func main() {
	l1 := createLinkedList([]int{2, 8, 7, 8, 1})
	l2 := createLinkedList([]int{2, 8, 7, 8})

	r := Tashizan(l1, l2)

	fmt.Println(r)
}
