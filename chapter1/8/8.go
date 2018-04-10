package main

import (
	"fmt"
	"strings"
)

func isRotatedString(s, rotated string) bool {
	if len(s) != len(rotated) {
		return false
	}
	ss := s + s
	return strings.Contains(ss, rotated)
}

func main() {
	s := "waterbottle"
	rotated := "totlewaterb"

	fmt.Println(isRotatedString(s, rotated))
}
