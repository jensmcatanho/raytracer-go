package camera

import (
	"errors"
	"fmt"
	"image"
	"jensmcatanho/raytracer-go/math"
)

// ProjectionPlane is the structure of a plane in which the 3D scene is projected onto
type ProjectionPlane struct {
	Width     int
	Height    int
	PixelSize float64
	Distance  float64
	Gamma     float64

	ClampOutOfGamut bool
	ClampColor      math.Color

	Image image.RGBA
}

// NewProjectionPlane creates a ProjectionPlane structure
func NewProjectionPlane(args ...interface{}) (*ProjectionPlane, error) {
	width, height, pixelSize, distance, err := projectionPlaneParams(args)

	projectionPlane := &ProjectionPlane{
		Width:           width,
		Height:          height,
		PixelSize:       pixelSize,
		Distance:        distance,
		Gamma:           1.0,
		ClampOutOfGamut: false,
		ClampColor:      *math.NewColor(0., 0., 0.),
		Image:           *image.NewRGBA(image.Rect(0, 0, width, height)),
	}

	return projectionPlane, err
}

func projectionPlaneParams(args []interface{}) (width, height int, pixelSize, distance float64, err error) {
	pixelSize = 1.
	distance = 500.

	if len(args) < 2 || len(args) > 4 {
		err = fmt.Errorf("Invalid number of arguments: %d arguments received", len(args))
		return
	}

	for i, paramInterface := range args {
		switch i {
		case 0:
			param, ok := paramInterface.(int)
			if !ok {
				err = errors.New("1st argument is not of type int")
				return
			}

			width = param

		case 1:
			param, ok := paramInterface.(int)
			if !ok {
				err = errors.New("2nd argument is not of type int")
				return
			}

			height = param

		case 2:
			param, ok := paramInterface.(float64)
			if !ok {
				err = errors.New("3rd argument is not of type float64")
				return
			}

			pixelSize = param

		case 3:
			param, ok := paramInterface.(float64)
			if !ok {
				err = errors.New("4th argument is not of type float64")
				return
			}

			distance = param
		}
	}

	return
}

// SetPixel sets a color to a pixel in the projection plane
func (p *ProjectionPlane) SetPixel(row, col int, rawColor math.Color) {
	mappedColor := rawColor

	if p.ClampOutOfGamut {
		mappedColor.ClampToColor(p.ClampColor)
	} else {
		mappedColor.MaxToOne()
	}

	if p.Gamma != 1.0 {
		mappedColor = *mappedColor.Pow(1 / p.Gamma)
	}

	p.Image.Set(row, col, mappedColor.ToRGBA())
}
