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

func (ll *LinkedList) Print() {
	n := ll.head
	for {
		if n == nil {
			return
		}

		fmt.Printf("%d", n.Value)
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

func DeleteNode(n *LinkedListNode) {
	nn := n.NextNode

	n.Value = nn.Value
	n.NextNode = nn.NextNode
}

func main() {
	ll := createLinkedList([]int{1, 2, 3})

	// 2
	c := ll.head.NextNode

	DeleteNode(c)

	ll.Print()
}
