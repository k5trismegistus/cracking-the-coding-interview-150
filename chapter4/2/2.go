package main

import (
	"fmt"
)

type AdjList []Adj

type Adj []int

func InSlice(sl []int, elem int) bool {
	for _, b := range sl {
		if b == elem {
			return true
		}
	}
	return false
}

func (al *AdjList) Walk(start int) bool {
	visited := []int{}
	routes := [][]int{[]int{start}}

	for {
		newRoutes := [][]int{}

		for _, route := range routes {
			lastVisited := route[len(route)-1]

			candidates := (*al)[lastVisited]

			for _, c := range candidates {
				if c == start {
					return true
				}

				if !InSlice(visited, c) {
					newRoute := append(route, c)
					newRoutes = append(newRoutes, newRoute)
				}
			}
		}

		if len(newRoutes) == 0 {
			break
		}

		routes = newRoutes
	}

	return false
}

func main() {
	al := &AdjList{
		Adj{1, 2},
		Adj{},
		Adj{3},
		Adj{0},
	}

	fmt.Println(al.Walk(0))
}
