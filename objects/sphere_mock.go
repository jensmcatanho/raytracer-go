package objects

import (
	"jensmcatanho/raytracer-go/math"

	"github.com/stretchr/testify/mock"
)

// SphereMock is a mock for the Sphere structure
type SphereMock struct {
	mock.Mock
}

// Hit is a mock for the Hit method
func (s *SphereMock) Hit(ray math.Ray, closestDistance *float64) *math.Surface {
	args := s.Called(ray, closestDistance)
	return args.Get(0).(*math.Surface)
}

// hitSurface is a mock for the hitSurface method
func (s *SphereMock) hitSurface(ray math.Ray, distance float64, originToCenter *math.Vector) *math.Surface {
	args := s.Called(ray, distance, originToCenter)
	return args.Get(0).(*math.Surface)
}
