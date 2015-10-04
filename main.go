package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	imagePath, _ := filepath.Abs(os.Args[1])
	width, _ := strconv.Atoi(os.Args[2])
	height, _ := strconv.Atoi(os.Args[3])

	imageFile, _ := os.Create(imagePath)
	makeFractal(imageFile, width, height)
}

func makeFractal(imageFile *os.File, width int, height int) {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	bounds := img.Bounds()

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c := color.RGBA{r(x, y), g(x, y), b(x, y), 255}
			img.Set(x, y, c)
		}
	}

	png.Encode(imageFile, img)
}

func r(x, y int) uint8 {
	return uint8(x ^ y - x ^ x)
}

func g(x, y int) uint8 {
	return uint8(x | y)
}

func b(x, y int) uint8 {
	return uint8(x ^ y - y ^ y)
}
