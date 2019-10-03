package main

import (
	"fmt"
	"jensmcatanho/raytracer-go/camera"
	"jensmcatanho/raytracer-go/math"
	"jensmcatanho/raytracer-go/object"
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
	sphere := object.NewSphere(*sphereCenter, 2)
	scene.GetInstance().AddObject(sphere)

	projectionPlane, err := camera.NewProjectionPlane(200, 200, 1., 500.)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error on projection plane creation: %s", err))
	}

	camera := new(camera.Pinhole)
	camera.ProjectionPlane = *projectionPlane
	camera.Eye = math.NewVector(.0, .0, .0)
	camera.LookAt = math.NewVector(.0, .0, 1.)
	camera.Exposure = 1
	camera.Zoom = 1
	camera.Tracer = new(tracer.RayCast)
	camera.RenderScene()

	fmt.Println(fmt.Sprintf("Raytracer Go!\nScene: %+v", scene.GetInstance()))
}
