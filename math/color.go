package math

import stdMath "math"

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
		r: stdMath.Pow(c.r, k),
		g: stdMath.Pow(c.g, k),
		b: stdMath.Pow(c.b, k),
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
	max := stdMath.Max(c.r, stdMath.Max(c.b, c.g))

	if max > 1.0 {
		c.r /= max
		c.g /= max
		c.b /= max
	}
}
