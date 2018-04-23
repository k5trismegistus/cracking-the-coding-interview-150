package main

import (
	"fmt"
)

type MyQueue struct {
	Data Stack
	Tmp  Stack
}

type Stack struct {
	Head *Node
}

type Node struct {
	Value int
	Next  *Node
}

func (s *Stack) Push(v int) {
	n := &Node{Value: v}

	n.Next = s.Head
	s.Head = n
}

func (s *Stack) Pop() int {
	n := s.Head
	s.Head = n.Next

	return n.Value
}

func (q *MyQueue) Enqueue(v int) {
	q.Data.Push(v)
}

func (q *MyQueue) Dequeue() int {
	for {
		if q.Data.Head == nil {
			break
		}

		v := q.Data.Pop()
		q.Tmp.Push(v)
	}

	r := q.Tmp.Pop()

	for {
		if q.Tmp.Head == nil {
			break
		}

		v := q.Tmp.Pop()
		q.Data.Push(v)
	}

	return r
}

func main() {
	q := &MyQueue{}

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	fmt.Println(q.Dequeue())
	q.Enqueue(4)
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	q.Enqueue(5)
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
}
