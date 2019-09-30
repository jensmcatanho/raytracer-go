package math

// Ray is a structure that represent an infinite straight line
type Ray struct {
	Origin    Vector
	Direction Vector
}

// NewRay creates a Ray structure
func NewRay(origin, direction Vector) *Ray {
	return &Ray{
		Origin:    origin,
		Direction: direction,
	}
}
