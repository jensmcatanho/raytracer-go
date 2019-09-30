package objects

import "jensmcatanho/raytracer-go/math"

// Renderable is an interface for objects that can be hit by rays
type Renderable interface {
	Hit(ray math.Ray, closestDistance *float64) *math.Surface
}
