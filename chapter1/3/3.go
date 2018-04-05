package main

import (
	"fmt"
)

func anagram(s1, s2 string) bool {
	m := make(map[rune]int32)

	for _, c := range s1 {
		m[c] = m[c] + 1
	}

	for _, c := range s2 {
		m[c] = m[c] - 1
		if m[c] < 0 {
			return false
		}
	}

	for _, v := range m {
		if v != 0 {
			return false
		}
	}

	return true
}

func main() {
	s1 := "abcdefq11"
	s2 := "dbcafqe1"

	b := anagram(s1, s2)
	fmt.Println(b)
}
