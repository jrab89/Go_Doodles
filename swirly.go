package main

import "math"

func r(x, y int) uint8 {
	top := math.Hypot(float64(148-x), float64(1000-y)) + 1
	bottom := math.Sqrt(float64(math.Abs(math.Sin((math.Sqrt(float64(sq(500-x)+sq(400-y))))/115.0)))) + 1
	return uint8(top / bottom)
}

func g(x, y int) uint8 {
	// top := math.Sqrt(float64((148-x)*(148-x)+(1000-y)*(1000-y))) + 1
	// bottom := math.Sqrt(float64(math.Abs(math.Sin((math.Sqrt(float64((500-x)*(500-x)+(400-y)*(400-y))))/115.0)))) + 1
	// return uint8(top / bottom)
	return 0
}

func b(x, y int) uint8 {
	return 0
}

func sq(n int) int {
	return n * n
}
