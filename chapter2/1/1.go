package main

import (
	"fmt"
)

func inList(l []int, v int) bool {
	for _, lv := range l {
		if lv == v {
			return true
		}
	}
	return false
}

func undupV1(l []int) []int {
	e := []int{}
	nl := []int{}

	for _, v := range l {
		if inList(e, v) {
			continue
		}

		e = append(e, v)
		nl = append(nl, v)
	}

	return nl
}

func undupV2(l []int) []int {
	nl := []int{}

	for i, v1 := range l {
		for j, v2 := range l {
			if i == 0 {
				break
			}

			if i <= j {
				goto LS
			}

			if v1 == v2 {
				goto LS
			}
			break
		}
		nl = append(nl, v1)
	LS:
		continue
	}
	return nl
}

func main() {
	l := []int{1, 2, 3, 4, 1, 5, 6, 7, 1}

	fmt.Println(undupV1(l))
	fmt.Println(undupV2(l))
}
