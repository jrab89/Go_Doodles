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
			opacity := color.Gray{uint8(x ^ y)}
			img.Set(x, y, opacity)
		}
	}

	png.Encode(imageFile, img)
}
