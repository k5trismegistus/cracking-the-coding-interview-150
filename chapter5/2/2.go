package main

import (
	"fmt"
)

func Bin(v float64) string {
	if v < 0 || v >= 1 {
		return "ERROR"
	}

	t := v
	r := "0."
	frac := 0.5

	for {
		if len(r) > 32 {
			break
		}

		if t >= frac {
			t = t - frac
			r += "1"
		} else {
			r += "0"
		}

		if t == 0 {
			return r
		}

		frac = frac / 2
	}

	return "ERROR"
}

func main() {
	v := 0.625
	fmt.Println(Bin(v))
}
