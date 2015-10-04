package main

import "math"

func r(x, y int) uint8 {
	return b(x-y, y-x)
}

func g(x, y int) uint8 {
	return b(300-x, 300-y)
}

func b(x, y int) uint8 {
	return uint8(math.Tan(float64(x^y))) / 2
}
