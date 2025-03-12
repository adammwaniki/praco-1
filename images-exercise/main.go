package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

// Define Image type
type Image struct {
	width, height int
}

// Implement the Bounds method
func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.width, img.height)
}

// Implement the ColorModel method
func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

// Implement the At method
func (img Image) At(x, y int) color.Color {
	v := uint8(x ^ y) // XOR function to generate patterns or (x+y)/2 or x*y
	return color.RGBA{v, v, 255, 255} // Blue-tinted pattern
}

func main() {
	m := Image{256, 256} // Create an instance of Image with dimensions 256x256
	pic.ShowImage(m)     // Display the generated image
}
