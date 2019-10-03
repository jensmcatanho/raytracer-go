package object

import (
	"jensmcatanho/raytracer-go/math"
	stdMath "math"
)

const (
	epsilon = 0.01
)

// Sphere is a structure that represents a 3D sphere
type Sphere struct {
	Center math.Vector
	Radius float64
	Color  math.Color
}

// NewSphere creates a Sphere structure
func NewSphere(center math.Vector, radius float64) *Sphere {
	return &Sphere{
		Center: center,
		Radius: radius,
	}
}

// Hit checks if a ray hits the sphere and returns the surface hit
func (s *Sphere) Hit(ray math.Ray, closestDistance *float64) *math.Surface {
	originToCenter := ray.Origin.Sub(&s.Center)

	a := ray.Direction.Dot(&ray.Direction)
	b := originToCenter.Multiply(2.).Dot(&ray.Direction)
	c := originToCenter.Dot(originToCenter) - s.Radius*s.Radius
	discriminant := b*b - (4 * a * c)

	if discriminant < 0. {
		return &math.Surface{
			Hit: false,
		}
	}

	distance := (-b - stdMath.Sqrt(discriminant)) / 2 * a
	if distance > epsilon {
		*closestDistance = distance
		return s.hitSurface(ray, distance, originToCenter)
	}

	distance = (-b + stdMath.Sqrt(discriminant)) / 2 * a
	if distance > epsilon {
		*closestDistance = distance
		return s.hitSurface(ray, distance, originToCenter)
	}

	return &math.Surface{
		Hit: false,
	}
}

func (s *Sphere) hitSurface(ray math.Ray, distance float64, originToCenter *math.Vector) *math.Surface {
	hitPoint := ray.Origin.Add(ray.Direction.Multiply(distance))
	normal := originToCenter.Add(ray.Direction.Multiply(distance))
	normal = normal.Multiply(1. / s.Radius)

	return math.NewSurface(s.Color, true, *hitPoint, *normal)
}
