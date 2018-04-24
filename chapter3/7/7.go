package main

import (
	"fmt"
)

type Queue struct {
	Head *Node
}

func (q *Queue) LastNode() *Node {
	n := q.Head
	for {
		if n.Next == nil {
			return n
		}
		n = n.Next
	}
}

func (q *Queue) Enqueue(n *Node) {
	if q.Head == nil {
		q.Head = n
		return
	}

	q.LastNode().Next = n
}

func (q *Queue) DequeueAny() *Node {
	n := q.Head
	if n == nil {
		return nil
	}

	q.Head = n.Next
	return n
}

func (q *Queue) DequeueDog() *Node {
	n := q.Head

	for {
		if n == nil {
			return nil
		}

		if n.IsDog {
			d := &Node{
				Name:  n.Name,
				IsDog: n.IsDog,
			}

			n.Name = n.Next.Name
			n.IsDog = n.Next.IsDog
			n.Next = n.Next.Next

			return d
		}

		n = n.Next
	}
}

func (q *Queue) DequeueCat() *Node {
	n := q.Head

	for {
		if n == nil {
			return nil
		}

		if !n.IsDog {
			c := &Node{
				Name:  n.Name,
				IsDog: n.IsDog,
			}

			n.Name = n.Next.Name
			n.IsDog = n.Next.IsDog
			n.Next = n.Next.Next

			return c
		}

		n = n.Next
	}
}

type Node struct {
	Name  string
	IsDog bool
	Next  *Node
}

func main() {
	animalShelter := &Queue{}

	animalShelter.Enqueue(&Node{Name: "dog1", IsDog: true})
	animalShelter.Enqueue(&Node{Name: "cat1", IsDog: false})
	animalShelter.Enqueue(&Node{Name: "dog2", IsDog: true})
	animalShelter.Enqueue(&Node{Name: "dog3", IsDog: true})
	animalShelter.Enqueue(&Node{Name: "cat2", IsDog: false})
	animalShelter.Enqueue(&Node{Name: "cat3", IsDog: false})
	animalShelter.Enqueue(&Node{Name: "cat4", IsDog: false})

	fmt.Println(animalShelter.DequeueCat().Name)
	fmt.Println(animalShelter.DequeueCat().Name)
	fmt.Println(animalShelter.DequeueAny().Name)
	fmt.Println(animalShelter.DequeueAny().Name)
	fmt.Println(animalShelter.DequeueCat().Name)
	fmt.Println(animalShelter.DequeueAny().Name)
}
