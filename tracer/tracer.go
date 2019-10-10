package tracer

import "jensmcatanho/raytracer-go/math/geometry"

// Tracer is a interface for tracing rays
type Tracer interface {
	TraceRay(ray geometry.Ray) *geometry.Surface
}
