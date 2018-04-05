package main

import (
	"fmt"
)

func reverse(s string) string {
	r := []rune{}

	for _, c := range s {
		r = append([]rune{c}, []rune(r)...)
	}
	return string(r)
}

func main() {
	s := "abcdef日"

	rs := reverse(s)
	fmt.Println(rs)
}
