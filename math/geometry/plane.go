package geometry

import "jensmcatanho/raytracer-go/math/color"

// Plane is a structure that representes a 3D Plane
type Plane struct {
	Point  Vector
	Normal Vector
	Color  color.Color
}

// NewPlane creates a Plane structure
func NewPlane(point, normal Vector) *Plane {
	normal.Normalize()

	return &Plane{
		Point:  point,
		Normal: normal,
	}
}

// Hit checks if a ray hits the plane and returns the surface hit
func (p *Plane) Hit(ray Ray, closestDistance *float64) *Surface {
	planeToOrigin := p.Point.Sub(&ray.Origin)
	distance := planeToOrigin.Dot(&p.Normal) / ray.Direction.Dot(&p.Normal)

	if distance > epsilon {
		*closestDistance = distance
		hitPoint := ray.Origin.Add(ray.Direction.Multiply(distance))
		return NewSurface(p.Color, true, *hitPoint, p.Normal)
	}

	return &Surface{
		Hit: false,
	}
}
