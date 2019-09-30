package math

// Surface is a structure that represent a surface that was hit by a ray
type Surface struct {
	Color    *Color
	Hit      bool
	HitPoint *Vector
	Normal   *Vector
}

// NewSurface creates a Surface structure
func NewSurface(color Color, hit bool, hitPoint, normal Vector) *Surface {
	if hit {
		return &Surface{
			Color:    &color,
			Hit:      hit,
			HitPoint: &hitPoint,
			Normal:   &normal,
		}
	}

	return &Surface{
		Color:    nil,
		Hit:      hit,
		HitPoint: nil,
		Normal:   nil,
	}
}
