package tracer

import (
	"jensmcatanho/raytracer-go/math/color"
	"jensmcatanho/raytracer-go/math/geometry"
	"jensmcatanho/raytracer-go/scene"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_RayCast_TraceRay_WhenThereIsNoObjectToHit(t *testing.T) {
	backgroundColor := color.NewColor(0., 0., 0.)
	err := scene.GetInstance().Initialize(*backgroundColor)
	assert.Nil(t, err)

	ray := geometry.NewRay(*geometry.NewVector(0., 0., 0.), *geometry.NewVector(0., 0., 1.))
	tracer := new(RayCast)
	surface := tracer.TraceRay(*ray)
	assert.Nil(t, surface)
}

func Test_RayCast_TraceRay_WhenNoObjectIsHit(t *testing.T) {
	backgroundColor := color.NewColor(0., 0., 0.)
	err := scene.GetInstance().Initialize(*backgroundColor)
	assert.Nil(t, err)

	sphereCenter := geometry.NewVector(0., 10., 00.)
	sphere := geometry.NewSphere(*sphereCenter, 2)
	scene.GetInstance().AddObject(sphere)

	ray := geometry.NewRay(*geometry.NewVector(0., 0., 0.), *geometry.NewVector(0., 0., 1.))
	tracer := new(RayCast)
	surface := tracer.TraceRay(*ray)
	assert.Nil(t, surface)
}

func Test_RayCast_TraceRay_WhenAnObjectIsHit(t *testing.T) {
	backgroundColor := color.NewColor(0., 0., 0.)
	err := scene.GetInstance().Initialize(*backgroundColor)
	assert.Nil(t, err)

	sphereCenter := geometry.NewVector(0., 0., 10.)
	sphere := geometry.NewSphere(*sphereCenter, 2)
	sphere.Color = *color.NewColor(1., 0., 0.)
	scene.GetInstance().AddObject(sphere)

	ray := geometry.NewRay(*geometry.NewVector(0., 0., 0.), *geometry.NewVector(0., 0., 1.))
	tracer := new(RayCast)
	surface := tracer.TraceRay(*ray)
	expectedSurface := geometry.NewSurface(
		*color.NewColor(1., 0., 0.),
		true,
		*geometry.NewVector(0., 0., 8.),
		*geometry.NewVector(0., 0., -1.),
	)
	assert.Equal(t, expectedSurface, surface)
}

func Test_RayCast_TraceRay_WhenTwoObjectsAreHitAndOneIsInFrontOfTheOther(t *testing.T) {
	backgroundColor := color.NewColor(0., 0., 0.)
	err := scene.GetInstance().Initialize(*backgroundColor)
	assert.Nil(t, err)

	sphereCenter := geometry.NewVector(0., 0., 10.)
	sphere := geometry.NewSphere(*sphereCenter, 2)
	sphere.Color = *color.NewColor(1., 0., 0.)
	scene.GetInstance().AddObject(sphere)

	furtherSphereCenter := geometry.NewVector(0., 0., 20.)
	furtherSphere := geometry.NewSphere(*furtherSphereCenter, 2)
	scene.GetInstance().AddObject(furtherSphere)

	ray := geometry.NewRay(*geometry.NewVector(0., 0., 0.), *geometry.NewVector(0., 0., 1.))
	tracer := new(RayCast)
	surface := tracer.TraceRay(*ray)
	expectedSurface := geometry.NewSurface(
		*color.NewColor(1., 0., 0.),
		true,
		*geometry.NewVector(0., 0., 8.),
		*geometry.NewVector(0., 0., -1.),
	)
	assert.Equal(t, expectedSurface, surface)
}
