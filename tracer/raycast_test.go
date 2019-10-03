package tracer

import (
	"jensmcatanho/raytracer-go/math"
	"jensmcatanho/raytracer-go/object"
	"jensmcatanho/raytracer-go/scene"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_RayCast_TraceRay_WhenThereIsNoObjectToHit(t *testing.T) {
	backgroundColor := math.NewColor(0., 0., 0.)
	err := scene.GetInstance().Initialize(*backgroundColor)
	assert.Nil(t, err)

	ray := math.NewRay(*math.NewVector(0., 0., 0.), *math.NewVector(0., 0., 1.))
	tracer := new(RayCast)
	surface := tracer.TraceRay(*ray)
	assert.Nil(t, surface)
}

func Test_RayCast_TraceRay_WhenNoObjectIsHit(t *testing.T) {
	backgroundColor := math.NewColor(0., 0., 0.)
	err := scene.GetInstance().Initialize(*backgroundColor)
	assert.Nil(t, err)

	sphereCenter := math.NewVector(0., 10., 00.)
	sphere := object.NewSphere(*sphereCenter, 2)
	scene.GetInstance().AddObject(sphere)

	ray := math.NewRay(*math.NewVector(0., 0., 0.), *math.NewVector(0., 0., 1.))
	tracer := new(RayCast)
	surface := tracer.TraceRay(*ray)
	assert.Nil(t, surface)
}

func Test_RayCast_TraceRay_WhenAnObjectIsHit(t *testing.T) {
	backgroundColor := math.NewColor(0., 0., 0.)
	err := scene.GetInstance().Initialize(*backgroundColor)
	assert.Nil(t, err)

	sphereCenter := math.NewVector(0., 0., 10.)
	sphere := object.NewSphere(*sphereCenter, 2)
	sphere.Color = *math.NewColor(1., 0., 0.)
	scene.GetInstance().AddObject(sphere)

	ray := math.NewRay(*math.NewVector(0., 0., 0.), *math.NewVector(0., 0., 1.))
	tracer := new(RayCast)
	surface := tracer.TraceRay(*ray)
	expectedSurface := math.NewSurface(
		*math.NewColor(1., 0., 0.),
		true,
		*math.NewVector(0., 0., 8.),
		*math.NewVector(0., 0., -1.),
	)
	assert.Equal(t, expectedSurface, surface)
}

func Test_RayCast_TraceRay_WhenTwoObjectsAreHitAndOneIsInFrontOfTheOther(t *testing.T) {
	backgroundColor := math.NewColor(0., 0., 0.)
	err := scene.GetInstance().Initialize(*backgroundColor)
	assert.Nil(t, err)

	sphereCenter := math.NewVector(0., 0., 10.)
	sphere := object.NewSphere(*sphereCenter, 2)
	sphere.Color = *math.NewColor(1., 0., 0.)
	scene.GetInstance().AddObject(sphere)

	furtherSphereCenter := math.NewVector(0., 0., 20.)
	furtherSphere := object.NewSphere(*furtherSphereCenter, 2)
	scene.GetInstance().AddObject(furtherSphere)

	ray := math.NewRay(*math.NewVector(0., 0., 0.), *math.NewVector(0., 0., 1.))
	tracer := new(RayCast)
	surface := tracer.TraceRay(*ray)
	expectedSurface := math.NewSurface(
		*math.NewColor(1., 0., 0.),
		true,
		*math.NewVector(0., 0., 8.),
		*math.NewVector(0., 0., -1.),
	)
	assert.Equal(t, expectedSurface, surface)
}
