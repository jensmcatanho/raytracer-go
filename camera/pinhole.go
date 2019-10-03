package camera

import (
	"jensmcatanho/raytracer-go/math"
	"jensmcatanho/raytracer-go/tracer"
)

// Pinhole is a camera that renders a scene with a perspective projection
type Pinhole struct {
	Eye    *math.Vector
	LookAt *math.Vector

	u *math.Vector
	v *math.Vector
	w *math.Vector

	Yaw   float64
	Pitch float64
	Roll  float64

	Exposure float64
	Zoom     float64

	ProjectionPlane ProjectionPlane
	Tracer          tracer.Tracer
}

// RenderScene traces a ray for every pixel in the target image and sets the pixel color with the color of the object hit
func (p *Pinhole) RenderScene() {
	ray := new(math.Ray)
	ray.Origin = *p.Eye
	pixelSize := p.ProjectionPlane.PixelSize / p.Zoom
	p.ComputeUVW()

	for row := 0; row < p.ProjectionPlane.Height; row++ {
		for col := 0; col < p.ProjectionPlane.Width; col++ {
			var pixelColor math.Color
			samplePoint := math.NewVector(pixelSize*.5, pixelSize*.5, .0)
			ray.Direction = p.RayDirection(*samplePoint)

			surface := p.Tracer.TraceRay(*ray)
			if surface.Hit {
				pixelColor = *surface.Color
			} else {
				pixelColor = *math.NewColor(0., 0., 0.)
			}

			pixelColor = *pixelColor.Multiply(p.Exposure)
			p.ProjectionPlane.SetPixel(row, col, pixelColor)
		}
	}
}

// ComputeUVW calculates the camera's coordinate system
func (p *Pinhole) ComputeUVW() {
	if p.Eye.X == p.LookAt.X && p.Eye.Z == p.LookAt.Z && p.Eye.Y > p.LookAt.Y {
		// Looking down
		p.u = math.NewVector(0., 0., 1.)
		p.v = math.NewVector(1., 0., 0.)
		p.w = math.NewVector(0., 1., 0.)
		return
	} else if p.Eye.X == p.LookAt.X && p.Eye.Z == p.LookAt.Z && p.Eye.Y < p.LookAt.Y {
		// Looking up
		p.u = math.NewVector(1., 0., 0.)
		p.v = math.NewVector(0., 0., 1.)
		p.w = math.NewVector(0., -1., 0.)
		return
	}

	p.w = p.Eye.Sub(p.LookAt)
	p.w.Normalize()

	p.u = worldUp.Cross(p.w)
	p.u.Normalize()

	p.v = p.w.Cross(p.u)
}

// RayDirection returns the direction of a ray given a sample point on the projection plane
func (p *Pinhole) RayDirection(samplePoint math.Vector) math.Vector {
	direction := p.u.Multiply(samplePoint.X)
	direction = direction.Add(p.v.Multiply(samplePoint.Y))
	direction = direction.Sub(p.w.Multiply(p.ProjectionPlane.Distance))
	direction.Normalize()

	return *direction
}
