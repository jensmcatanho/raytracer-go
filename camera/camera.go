package camera

import "jensmcatanho/raytracer-go/math/geometry"

var (
	worldUp = geometry.NewVector(0., 1., 0.)
)

// Camera is an interface for structures that can render a scene
type Camera interface {
	RenderScene()
	RayDirection(samplePoint geometry.Vector) geometry.Vector
	SaveImage()
}
