package tracer

import "jensmcatanho/raytracer-go/math"

// Tracer is a interface for tracing rays
type Tracer interface {
	TraceRay(ray math.Ray) *math.Surface
}
