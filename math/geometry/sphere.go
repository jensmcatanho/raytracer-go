package geometry

import (
	"jensmcatanho/raytracer-go/math/color"
	"math"
)

// Sphere is a structure that represents a 3D sphere
type Sphere struct {
	Center Vector
	Radius float64
	Color  color.Color
}

// NewSphere creates a Sphere structure
func NewSphere(center Vector, radius float64) *Sphere {
	return &Sphere{
		Center: center,
		Radius: radius,
	}
}

// Hit checks if a ray hits the sphere and returns the surface hit
func (s *Sphere) Hit(ray Ray, closestDistance *float64) *Surface {
	originToCenter := ray.Origin.Sub(&s.Center)

	a := ray.Direction.Dot(&ray.Direction)
	b := originToCenter.Multiply(2.).Dot(&ray.Direction)
	c := originToCenter.Dot(originToCenter) - s.Radius*s.Radius
	discriminant := b*b - (4 * a * c)

	if discriminant < 0. {
		return &Surface{
			Hit: false,
		}
	}

	distance := (-b - math.Sqrt(discriminant)) / 2 * a
	if distance > epsilon {
		*closestDistance = distance
		return s.hitSurface(ray, distance, originToCenter)
	}

	distance = (-b + math.Sqrt(discriminant)) / 2 * a
	if distance > epsilon {
		*closestDistance = distance
		return s.hitSurface(ray, distance, originToCenter)
	}

	return &Surface{
		Hit: false,
	}
}

func (s *Sphere) hitSurface(ray Ray, distance float64, originToCenter *Vector) *Surface {
	hitPoint := ray.Origin.Add(ray.Direction.Multiply(distance))
	normal := originToCenter.Add(ray.Direction.Multiply(distance))
	normal = normal.Multiply(1. / s.Radius)

	return NewSurface(s.Color, true, *hitPoint, *normal)
}
