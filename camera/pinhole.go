package camera

import (
	"image/png"
	"jensmcatanho/raytracer-go/math/color"
	"jensmcatanho/raytracer-go/math/geometry"
	"jensmcatanho/raytracer-go/math/sampler"
	"jensmcatanho/raytracer-go/tracer"
	"log"
	"os"
)

// Pinhole is a camera that renders a scene with a perspective projection
type Pinhole struct {
	Eye    *geometry.Vector
	LookAt *geometry.Vector

	u *geometry.Vector
	v *geometry.Vector
	w *geometry.Vector

	Yaw   float64
	Pitch float64
	Roll  float64

	Exposure float64
	Zoom     float64

	ProjectionPlane ProjectionPlane
	Tracer          tracer.Tracer
	Sampler         sampler.Sampler
}

// RenderScene traces a ray for every pixel in the target image and sets the pixel color with the color of the object hit
func (p *Pinhole) RenderScene() {
	ray := new(geometry.Ray)
	ray.Origin = *p.Eye
	pixelSize := p.ProjectionPlane.PixelSize / p.Zoom
	p.ComputeUVW()

	for row := 0; row < p.ProjectionPlane.Height; row++ {
		for col := 0; col < p.ProjectionPlane.Width; col++ {
			pixelColor := *color.NewColor(0., 0., 0.)

			for sample := 0; sample < p.Sampler.Samples; sample++ {
				samplePoint := p.Sampler.SampleUnitSquare()
				ray.Direction = p.RayDirection(*geometry.NewVector(
					pixelSize*(float64(col)-0.5*float64(p.ProjectionPlane.Width)+samplePoint.X),
					pixelSize*(float64(row)-0.5*float64(p.ProjectionPlane.Height)+samplePoint.Y),
					0.,
				))

				surface := p.Tracer.TraceRay(*ray)
				if surface != nil && surface.Hit {
					pixelColor = *pixelColor.Add(surface.Color)
				}
			}

			pixelColor = *pixelColor.Multiply(p.Exposure / float64(p.Sampler.Samples))
			p.ProjectionPlane.SetPixel(row, col, pixelColor)
		}
	}
}

// ComputeUVW calculates the camera's coordinate system
func (p *Pinhole) ComputeUVW() {
	if p.Eye.X == p.LookAt.X && p.Eye.Z == p.LookAt.Z && p.Eye.Y > p.LookAt.Y {
		// Looking down
		p.u = geometry.NewVector(0., 0., 1.)
		p.v = geometry.NewVector(1., 0., 0.)
		p.w = geometry.NewVector(0., 1., 0.)
		return
	} else if p.Eye.X == p.LookAt.X && p.Eye.Z == p.LookAt.Z && p.Eye.Y < p.LookAt.Y {
		// Looking up
		p.u = geometry.NewVector(1., 0., 0.)
		p.v = geometry.NewVector(0., 0., 1.)
		p.w = geometry.NewVector(0., -1., 0.)
		return
	}

	p.w = p.Eye.Sub(p.LookAt)
	p.w.Normalize()

	p.u = worldUp.Cross(p.w)
	p.u.Normalize()

	p.v = p.w.Cross(p.u)
}

// RayDirection returns the direction of a ray given a sample point on the projection plane
func (p *Pinhole) RayDirection(samplePoint geometry.Vector) geometry.Vector {
	direction := p.u.Multiply(samplePoint.X)
	direction = direction.Add(p.v.Multiply(samplePoint.Y))
	direction = direction.Sub(p.w.Multiply(p.ProjectionPlane.Distance))
	direction.Normalize()

	return *direction
}

// SaveImage saves a png image of the rendered scene
func (p *Pinhole) SaveImage() {
	file, err := os.Create("image.png")
	if err != nil {
		log.Fatal(err)
	}

	if err := png.Encode(file, &p.ProjectionPlane.Image); err != nil {
		file.Close()
		log.Fatal(err)
	}

	if err := file.Close(); err != nil {
		log.Fatal(err)
	}
}
