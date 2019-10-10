package geometry

const (
	epsilon = 0.01
)

// Renderable is an interface for objects that can be hit by rays
type Renderable interface {
	Hit(ray Ray, closestDistance *float64) *Surface
}
