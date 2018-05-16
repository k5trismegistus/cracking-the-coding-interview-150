package main

import (
	"fmt"
	"strconv"
)

func Set(b int, i uint) int {
	return b | (1 << i)
}

func Clear(b int, i uint) int {
	mask := ^(1 << i)
	return b & mask
}

func Update(b int, i uint, v int) int {
	r := Clear(b, i)
	if v != 0 {
		r = Set(b, i)
	}
	return r
}

func main() {
	b := "1000000000"
	r64, _ := strconv.ParseInt(b, 2, 0)
	r := int(r64)

	i := 2
	j := 6
	m := "10011"
	mr, _ := (strconv.ParseInt(m, 2, 0))

	for idx := i; idx < j; idx++ {
		r = Clear(r, uint(idx))
	}

	r = (r | (int(mr) << uint(i)))
	fmt.Println(strconv.FormatInt(int64(r), 2))
}
