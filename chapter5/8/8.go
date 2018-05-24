package main

import (
	"fmt"
)

func initScreen(width, height int) []byte {
	if width%8 != 0 {
		panic("width must be product of 8")
	}

	r := []byte{}

	for i := 0; i < width*height/8; i++ {
		r = append(r, 0x00)
	}

	return r
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func printScreen(screen []byte, width int) {
	line := ""
	for _, b := range screen {
		line += Reverse(fmt.Sprintf("%08b", b))
		if len(line) == width {
			fmt.Println(line)
			line = ""
		}
	}
}

func drawHorizontalLine(screen []byte, width, x1, x2, y int) {

	startCellIndex := (width*y + x1) / 8
	endCellIndex := (width*y + x2) / 8

	for i := startCellIndex + 1; i < endCellIndex; i++ {
		screen[i] = ^screen[i]
	}

	startOffset := 8 - (width*y+x1)%8
	if startOffset == 0 {
		startOffset = 8
	}

	for i := 0; i < 8-(width*y+x1)%8; i++ {
		screen[startCellIndex] = screen[startCellIndex] >> 1
		screen[startCellIndex] = (screen[startCellIndex] | 0x80)
	}

	for i := 0; i < (width*y+x2)%8; i++ {
		screen[endCellIndex] = screen[endCellIndex] << 1
		screen[endCellIndex] = (screen[endCellIndex] | 0x01)
	}
}

func main() {
	screen := initScreen(32, 32)
	printScreen(screen, 32)
	fmt.Println("=========================")

	drawHorizontalLine(screen, 32, 2, 31, 31)
	printScreen(screen, 32)
}
