package main

func r(x, y int) uint8 {
	return uint8((x+y)&y) / 2
}

func g(x, y int) uint8 {
	return uint8(x+y^y) / 2
}

func b(x, y int) uint8 {
	return uint8((255 + x - y) & y)
}
