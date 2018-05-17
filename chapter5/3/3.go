package main

import (
	"fmt"
	"strconv"
)

func smaller(v int64) int64 {
	b := []byte(strconv.FormatInt(int64(v), 2))

	for i := len(b) - 1; i > 1; i-- {
		if b[i-1] == byte('1') && b[i] == byte('0') {
			b[i-1] = byte('0')
			b[i] = byte('1')

			r, _ := strconv.ParseInt(string(b), 2, 64)
			return r
		}
	}

	return v
}

func bigger(v int64) int64 {
	b := []byte(strconv.FormatInt(int64(v), 2))

	for i := len(b) - 1; i > 1; i-- {
		if b[i-1] == byte('0') && b[i] == byte('1') {
			b[i-1] = byte('1')
			b[i] = byte('0')

			r, _ := strconv.ParseInt(string(b), 2, 64)
			return r
		}
	}

	return v
}

func main() {
	fmt.Println(strconv.FormatInt(10, 2))
	fmt.Println(strconv.FormatInt(smaller(10), 2))
	fmt.Println(strconv.FormatInt(bigger(10), 2))
}
