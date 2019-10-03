package camera

import "jensmcatanho/raytracer-go/math"

var (
	worldUp = math.NewVector(0., 1., 0.)
)

// Camera is an interface for structures that can render a scene
type Camera interface {
	RenderScene()
	RayDirection(samplePoint math.Vector) math.Vector
}
