package main

import (
	"fmt"
	"jensmcatanho/raytracer-go/camera"
	"jensmcatanho/raytracer-go/math"
	"jensmcatanho/raytracer-go/object"
	"jensmcatanho/raytracer-go/sampler"
	"jensmcatanho/raytracer-go/scene"
	"jensmcatanho/raytracer-go/tracer"
)

func main() {
	backgroundColor := math.NewColor(0., 0., 0.)
	err := scene.GetInstance().Initialize(*backgroundColor)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error on scene initializing: %s", err))
	}

	sphereCenter := math.NewVector(0., 0., 10.)
	sphere := object.NewSphere(*sphereCenter, 8)
	sphere.Color = *math.NewColor(1., 0., 0.)
	scene.GetInstance().AddObject(sphere)

	projectionPlane, err := camera.NewProjectionPlane(1000, 1000, 1., 100.)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error on projection plane creation: %s", err))
	}

	sampler, err := sampler.NewSampler(4, 1, sampler.Regular)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error on sampler creation: %s", err))
	}
	sampler.Sample()

	camera := new(camera.Pinhole)
	camera.ProjectionPlane = *projectionPlane
	camera.Sampler = *sampler
	camera.Eye = math.NewVector(.0, .0, .0)
	camera.LookAt = math.NewVector(.0, .0, 1.)
	camera.Exposure = 1
	camera.Zoom = 1
	camera.Tracer = new(tracer.RayCast)
	camera.RenderScene()
	camera.SaveImage()

	fmt.Println(fmt.Sprintf("Raytracer Go!\nScene: %+v", scene.GetInstance()))
}
