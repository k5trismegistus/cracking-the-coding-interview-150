package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Distance(x, y int64) int {
	bits := ^(x & y) & (x | y)
	s := strconv.FormatInt(bits, 2)

	return strings.Count(s, "1")
}

func main() {
	x := int64(31)
	y := int64(14)

	fmt.Println(Distance(x, y))
}
