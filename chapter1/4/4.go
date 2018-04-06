package main

import (
	"fmt"
)

func replace(s string) string {
	r := []rune{}

	for _, c := range s {
		if c == ' ' {
			r = append(r, '%', '2', '0')
			continue
		}
		r = append(r, c)
	}

	return string(r)
}

func main() {
	s := "Mr John Smith"
	fmt.Println(replace(s))
}
