package main

import "math"

func r(x, y int) uint8 {
	i := float64(x)
	j := float64(y)

	top := math.Hypot(148-i, 1000-j) + 1
	bottom := math.Sqrt(float64(math.Abs(math.Sin((math.Hypot(500-i, 400-j))/115)))) + 1
	return uint8(top / bottom)
}

func g(x, y int) uint8 {
	i := float64(x)
	j := float64(y)

	top := math.Hypot(610-i, 60-j) + 1
	bottom := math.Sqrt(float64(math.Abs(math.Sin((math.Hypot(864-i, 860-j))/115)))) + 1
	return uint8(top / bottom)
}

func b(x, y int) uint8 {
	i := float64(x)
	j := float64(y)

	top := math.Hypot(300-i, 20-j) + 1
	bottom := math.Sqrt(float64(math.Abs(math.Sin((math.Hypot(503-i, 103-j))/115)))) + 1
	return uint8(top / bottom)
}
