package geometry

import (
	"jensmcatanho/raytracer-go/math/color"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewSurface_WhenANewSurfaceIsCreated(t *testing.T) {
	color := color.NewColor(1., 1., 1.)
	hit := true
	hitPoint := NewVector(0., 0., 0.)
	normal := NewVector(0., 1., 0.)

	surface := NewSurface(*color, hit, *hitPoint, *normal)

	assert.Equal(t, color, surface.Color)
	assert.Equal(t, hit, surface.Hit)
	assert.Equal(t, hitPoint, surface.HitPoint)
	assert.Equal(t, normal, surface.Normal)
}

func Test_NewSurface_WhenNoSurfaceWasHit(t *testing.T) {
	color := color.NewColor(1., 1., 1.)
	hit := false
	hitPoint := NewVector(0., 0., 0.)
	normal := NewVector(0., 1., 0.)

	surface := NewSurface(*color, hit, *hitPoint, *normal)

	assert.Nil(t, surface.Color)
	assert.Equal(t, hit, surface.Hit)
	assert.Nil(t, surface.HitPoint)
	assert.Nil(t, surface.Normal)

}
