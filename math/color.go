package math

// Color is a structure that represents an RGB color
type Color struct {
	r, g, b float32
}

// NewColor creates a Color structure
func NewColor(r, g, b float32) *Color {
	return &Color{
		r: r,
		g: g,
		b: b,
	}
}
