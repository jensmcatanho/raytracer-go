package color

import (
	"image/color"
	"math"
)

// Color is a structure that represents an RGB color
type Color struct {
	r, g, b float64
}

// NewColor creates a Color structure
func NewColor(r, g, b float64) *Color {
	return &Color{
		r: r,
		g: g,
		b: b,
	}
}

// Add adds two colors and returns a new color with the result
func (c *Color) Add(d *Color) *Color {
	return &Color{
		r: c.r + d.r,
		g: c.g + d.g,
		b: c.b + d.b,
	}
}

// Multiply multiplies a Color by a constant and returns a new Color with the result
func (c *Color) Multiply(k float64) *Color {
	return &Color{
		r: c.r * k,
		g: c.g * k,
		b: c.b * k,
	}
}

func (c *Color) Pow(k float64) *Color {
	return &Color{
		r: math.Pow(c.r, k),
		g: math.Pow(c.g, k),
		b: math.Pow(c.b, k),
	}
}

func (c *Color) ClampToColor(target Color) {
	if c.r > 1.0 || c.g > 1.0 || c.b > 1.0 {
		c.r = target.r
		c.g = target.g
		c.b = target.b
	}
}

func (c *Color) MaxToOne() {
	max := math.Max(c.r, math.Max(c.b, c.g))

	if max > 1.0 {
		c.r /= max
		c.g /= max
		c.b /= max
	}
}

// ToRGBA converts float64 color fields to uint8
func (c *Color) ToRGBA() *color.RGBA {
	return &color.RGBA{
		R: uint8(c.r * 255),
		G: uint8(c.g * 255),
		B: uint8(c.b * 255),
		A: 255,
	}
}
