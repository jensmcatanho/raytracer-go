package tracer

import (
	"jensmcatanho/raytracer-go/math/geometry"
	"jensmcatanho/raytracer-go/scene"
	"math"
)

// RayCast is a structure of a tracer that casts rays
type RayCast struct {
}

// TraceRay iterates through the objects in the scene e checks if the ray intersects them
func (t *RayCast) TraceRay(ray geometry.Ray) *geometry.Surface {
	var closestSurface geometry.Surface
	var distance float64
	closestDistance := math.Inf(1)

	for _, object := range scene.GetInstance().ObjectList() {
		surface := *object.Hit(ray, &distance)

		if surface.Hit && distance < closestDistance {
			closestSurface = surface
			closestDistance = distance
		}
	}

	if closestSurface.Hit {
		return &closestSurface
	}

	return nil
}
