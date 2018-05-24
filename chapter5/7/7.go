package main

import (
	"fmt"
)

func PowInt(v, p int) int {
	r := 1
	for i := 0; i < p; i++ {
		r = r * v
	}
	return r
}

func Int2Boolist(v, size int) []bool {
	if v > PowInt(2, size) {
		panic("size is too small")
	}

	r := []bool{}

	for p := 1; p <= size; p++ {
		if v >= PowInt(2, size-p) {
			r = append(r, true)
			v = v - PowInt(2, size-p)
		} else {
			r = append(r, false)
		}
	}

	return r
}

func FindMissing(inputs [][]bool, idx int) int {
	if idx < 0 {
		return 0
	}

	lastOne := [][]bool{}
	lastZero := [][]bool{}

	for _, input := range inputs {
		if input[idx] == false {
			lastZero = append(lastZero, input)
		} else {
			lastOne = append(lastOne, input)
		}
	}

	if len(lastZero) <= len(lastOne) {
		v := FindMissing(lastZero, idx-1)
		return (v << 1) | 0
	} else {
		v := FindMissing(lastOne, idx-1)
		return (v << 1) | 1
	}
}

func main() {
	i := [][]bool{
		Int2Boolist(0, 3),
		Int2Boolist(1, 3),
		Int2Boolist(2, 3),
		Int2Boolist(3, 3),
		Int2Boolist(4, 3),
		// Int2Boolist(5, 3),
		Int2Boolist(6, 3),
		Int2Boolist(7, 3),
	}

	fmt.Println(FindMissing(i, 2))
}
