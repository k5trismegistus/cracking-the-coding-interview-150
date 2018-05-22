package main

import (
	"fmt"
)

func Swap(x uint64) uint64 {
	return ((x & 0xaaaaaaaaaaaaaaaa >> 1) | (x & 0x5555555555555555 << 1))
}

func main() {
	x := 5
	fmt.Println(Swap(uint64(x)))
}
