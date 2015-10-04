package main

func r(x, y int) uint8 {
	return uint8(x ^ y - x ^ x)
}

func g(x, y int) uint8 {
	return uint8(x | y)
}

func b(x, y int) uint8 {
	return uint8(x ^ y - y ^ y)
}
