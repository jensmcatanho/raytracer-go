package camera

import (
	"jensmcatanho/raytracer-go/math/geometry"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	epsilon = 0.001
)

func Test_Pinhole_ComputeUVW_WhenCameraIsLookingDown(t *testing.T) {
	camera := new(Pinhole)
	camera.Eye = geometry.NewVector(.0, .0, .0)
	camera.LookAt = geometry.NewVector(.0, -1., 0.)

	camera.ComputeUVW()
	assert.Equal(t, geometry.NewVector(0., 0., 1.), camera.u)
	assert.Equal(t, geometry.NewVector(1., 0., 0.), camera.v)
	assert.Equal(t, geometry.NewVector(0., 1., 0.), camera.w)
}

func Test_Pinhole_ComputeUVW_WhenCameraIsLookingUp(t *testing.T) {
	camera := new(Pinhole)
	camera.Eye = geometry.NewVector(.0, .0, .0)
	camera.LookAt = geometry.NewVector(.0, 1., 0.)

	camera.ComputeUVW()
	assert.Equal(t, geometry.NewVector(1., 0., 0.), camera.u)
	assert.Equal(t, geometry.NewVector(0., 0., 1.), camera.v)
	assert.Equal(t, geometry.NewVector(0., -1., 0.), camera.w)
}

func Test_Pinhole_ComputeUVW_WhenCameraIsLookingAnywhere(t *testing.T) {
	camera := new(Pinhole)
	camera.Eye = geometry.NewVector(.0, .0, .0)
	camera.LookAt = geometry.NewVector(.0, 0., 1.)

	camera.ComputeUVW()
	assert.Equal(t, geometry.NewVector(-1., 0., 0.), camera.u)
	assert.Equal(t, geometry.NewVector(0., -1., 0.), camera.v)
	assert.Equal(t, geometry.NewVector(0., 0., -1.), camera.w)
}

func Test_Pinhole_RayDirection(t *testing.T) {
	projectionPlane, err := NewProjectionPlane(200, 200, 1., 10.)
	assert.Nil(t, err)

	camera := new(Pinhole)
	camera.ProjectionPlane = *projectionPlane
	camera.Eye = geometry.NewVector(.0, .0, .0)
	camera.LookAt = geometry.NewVector(.0, .0, 1.)
	camera.ComputeUVW()

	samplePoint := geometry.NewVector(.5, .5, .0)
	direction := camera.RayDirection(*samplePoint)

	assert.LessOrEqual(t, math.Abs(-0.049875466805381644-direction.X), epsilon)
	assert.LessOrEqual(t, math.Abs(-0.049875466805381644-direction.Y), epsilon)
	assert.LessOrEqual(t, math.Abs(0.9975093361076328-direction.Z), epsilon)
}
