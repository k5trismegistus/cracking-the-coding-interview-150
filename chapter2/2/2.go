package main

import (
	"fmt"
)

func findLast(l []int, k int) int {
	if len(l) < k {
		panic("k is too big")
	}

	ll := len(l)
	var r int

	for i := range l {
		if i <= k {
			continue
		}

		if i == ll {
			break
		}
		r = l[i-k]
	}
	return r
}

func main() {
	l := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	k := 9

	fmt.Println(findLast(l, k))
}
