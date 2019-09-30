package objects

import (
	"jensmcatanho/raytracer-go/math"
	stdMath "math"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	infinity = stdMath.Inf(1)
)

func Test_NewSphere_WhenANewSphereIsCreated(t *testing.T) {
	center := math.NewVector(0., 0., 0.)
	radius := 10.

	sphere := NewSphere(*center, radius)
	assert.Equal(t, *center, sphere.Center)
	assert.Equal(t, radius, sphere.Radius)
}

func Test_Hit_WhenTheSphereIsNotHitAndRayOriginIsOutsideTheSphere(t *testing.T) {
	center := math.NewVector(0., 0., 0.)
	radius := 10.
	sphere := NewSphere(*center, radius)

	origin := math.NewVector(10., 10., 0.)
	direction := math.NewVector(0., 0., 1.)
	ray := math.NewRay(*origin, *direction)

	closestDistance := infinity
	surface := sphere.Hit(*ray, &closestDistance)
	assert.False(t, surface.Hit)
	assert.Equal(t, infinity, closestDistance)
}

func Test_Hit_WhenTheSphereIsNotHitAndRayOriginIsOnSphereSurface(t *testing.T) {
	center := math.NewVector(0., 0., 0.)
	radius := 10.
	sphere := NewSphere(*center, radius)

	origin := math.NewVector(10., 0., 0.)
	direction := math.NewVector(1., 0., 0.)
	ray := math.NewRay(*origin, *direction)

	closestDistance := infinity
	surface := sphere.Hit(*ray, &closestDistance)
	assert.False(t, surface.Hit)
	assert.Equal(t, infinity, closestDistance)
}

func Test_Hit_WhenTheSphereIsHitOnceAndRayOriginIsInsideSphere(t *testing.T) {
	center := math.NewVector(0., 0., 0.)
	radius := 10.
	sphere := NewSphere(*center, radius)

	origin := math.NewVector(0., 0., 0.)
	direction := math.NewVector(1., 0., 0.)
	ray := math.NewRay(*origin, *direction)

	closestDistance := infinity
	surface := sphere.Hit(*ray, &closestDistance)
	assert.True(t, surface.Hit)
	assert.Equal(t, 10., closestDistance)
}

func Test_Hit_WhenTheSphereIsHitOnceAndRayOriginIsOnSphereSurface(t *testing.T) {
	center := math.NewVector(0., 0., 0.)
	radius := 10.
	sphere := NewSphere(*center, radius)

	origin := math.NewVector(-10., 0., 0.)
	direction := math.NewVector(1., 0., 0.)
	ray := math.NewRay(*origin, *direction)

	closestDistance := infinity
	surface := sphere.Hit(*ray, &closestDistance)
	assert.True(t, surface.Hit)
	assert.Equal(t, 20., closestDistance)
}

func Test_Hit_When_TheSphereIsHitTwice(t *testing.T) {
	center := math.NewVector(0., 0., 0.)
	radius := 10.
	sphere := NewSphere(*center, radius)

	origin := math.NewVector(-15., 0., 0.)
	direction := math.NewVector(1., 0., 0.)
	ray := math.NewRay(*origin, *direction)

	closestDistance := infinity
	surface := sphere.Hit(*ray, &closestDistance)
	assert.True(t, surface.Hit)
	assert.Equal(t, 5., closestDistance)
}
