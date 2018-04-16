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

func (ll *LinkedList) Unshift(nn *LinkedListNode) {
	nn.NextNode = ll.head
	ll.head = nn
}

func (ll *LinkedList) Print() {
	n := ll.head

	for {
		if n == nil {
			return
		}

		fmt.Printf("%d, ", n.Value)
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

func (_ *LinkedList) Pop(n *LinkedListNode) *LinkedListNode {
	rn := &LinkedListNode{}
	rn.Value = n.Value

	nn := n.NextNode

	if nn != nil {
		n.Value = nn.Value
		n.NextNode = nn.NextNode
	} else {
		// Even if nil is assigned here, original value is shown...
		n = nil
	}

	return rn
}

func (ll *LinkedList) Mottekuru() {
	c := ll.head
	n := ll.head

	for {
		next := n.NextNode
		if n.Value < c.Value {
			rn := ll.Pop(n)
			ll.Unshift(rn)
		}

		if next == nil {
			break
		}

		n = next
	}
}

func main() {
	ll := createLinkedList([]int{50, 30, 55, 43})
	ll.Mottekuru()

	ll.Print()
}
