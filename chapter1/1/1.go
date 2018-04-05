package main

import (
	"fmt"
)

func unique(s string) bool {
	m := make(map[rune]int32)

	for _, c := range s {
		_, b := m[c]
		if b {
			return false
		}

		m[c] = 1
	}
	return true
}

func main() {
	s := "abcdef日"

	b := unique(s)
	fmt.Println(b)
}
