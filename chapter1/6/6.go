package main

import (
	"fmt"
)

// 正方行列なのは仮定してしまう
type Image [][][4]byte

func rotate(img Image, n int) Image {
	ri := make(Image, n)
	for i := range ri {
		ri[i] = make([][4]byte, n)
	}

	for i, l := range img {
		for j, v := range l {
			ri[j][n-i-1] = v
		}
	}

	return ri
}

func printImage(i Image) {
	for _, l := range i {
		fmt.Println(l)
	}
}

func main() {
	i := Image{
		{{' ', ' ', '0', '0'}, {' ', ' ', '0', '1'}, {' ', ' ', '0', '2'}, {' ', ' ', '0', '3'}},
		{{' ', ' ', '1', '0'}, {' ', ' ', '1', '1'}, {' ', ' ', '1', '2'}, {' ', ' ', '1', '3'}},
		{{' ', ' ', '2', '0'}, {' ', ' ', '2', '1'}, {' ', ' ', '2', '2'}, {' ', ' ', '2', '3'}},
		{{' ', ' ', '3', '0'}, {' ', ' ', '3', '1'}, {' ', ' ', '3', '2'}, {' ', ' ', '3', '3'}},
	}
	printImage(i)

	printImage(rotate(i, len(i)))
}
