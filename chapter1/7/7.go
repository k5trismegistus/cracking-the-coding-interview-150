package main

import (
	"fmt"
)

type Matrix [][]int

func inArray(arr []int, val int) bool {
	for _, b := range arr {
		if b == val {
			return true
		}
	}
	return false
}

func zeroFill(mtx Matrix) Matrix {
	n := len(mtx[0])

	toFillRow := []int{}
	toFillColumn := []int{}

	for i, row := range mtx {
		for j, value := range row {
			if value == 0 {
				toFillRow = append(toFillRow, i)
				toFillColumn = append(toFillColumn, j)
			}
		}
	}

	for _, i := range toFillRow {
		mtx[i] = make([]int, n)
	}

	for _, j := range toFillColumn {
		for i := range mtx {
			mtx[i][j] = 0
		}
	}

	return mtx
}

func printMatrix(mtx Matrix) {
	for _, l := range mtx {
		fmt.Println(l)
	}
}

func main() {
	mtx := [][]int{
		[]int{1, 0, 1, 1},
		[]int{1, 1, 1, 1},
		[]int{1, 1, 0, 1},
	}
	printMatrix(mtx)
	printMatrix(zeroFill(mtx))
}
