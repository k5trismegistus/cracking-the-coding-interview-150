package main

import (
	"fmt"
	"strconv"
)

func compress(s string) string {
	buf := []rune{}
	var cur rune
	var cnt int

	for _, c := range s {
		if cur == 0 {
			cur = c
			cnt = 1
			continue
		}

		if c != cur {
			sn := []rune(strconv.Itoa(cnt))
			buf = append(buf, cur)
			buf = append(buf, sn...)

			cur = c
			cnt = 1
			continue
		}

		cnt++
	}

	sn := []rune(strconv.Itoa(cnt))
	buf = append(buf, cur)
	buf = append(buf, sn...)

	compressed := string(buf)

	if len(compressed) < len(s) {
		return compressed
	}
	return s
}

func main() {
	s := "aabbccddeee"
	fmt.Println(compress(s))
}
