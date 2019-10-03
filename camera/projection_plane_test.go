package camera

import (
	"fmt"
	"jensmcatanho/raytracer-go/math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewProjectionPlane_WhenMinimumParamsArePassed(t *testing.T) {
	projectionPlane, err := NewProjectionPlane(800, 600)
	assert.Nil(t, err)
	assert.Equal(t, 800, projectionPlane.Width)
	assert.Equal(t, 600, projectionPlane.Height)
	assert.Equal(t, 1., projectionPlane.PixelSize)
	assert.Equal(t, 500., projectionPlane.Distance)
	assert.Equal(t, 1., projectionPlane.Gamma)
	assert.Equal(t, false, projectionPlane.ClampOutOfGamut)
	assert.Equal(t, *math.NewColor(0., 0., 0.), projectionPlane.ClampColor)
}

func Test_NewProjectionPlane_WhenPixelSizeIsPassed(t *testing.T) {
	projectionPlane, err := NewProjectionPlane(800, 600, 2.)
	assert.Nil(t, err)
	assert.Equal(t, 800, projectionPlane.Width)
	assert.Equal(t, 600, projectionPlane.Height)
	assert.Equal(t, 2., projectionPlane.PixelSize)
	assert.Equal(t, 500., projectionPlane.Distance)
	assert.Equal(t, 1., projectionPlane.Gamma)
	assert.Equal(t, false, projectionPlane.ClampOutOfGamut)
	assert.Equal(t, *math.NewColor(0., 0., 0.), projectionPlane.ClampColor)
}

func Test_NewProjectionPlane_WhenDistanceIsPassed(t *testing.T) {
	projectionPlane, err := NewProjectionPlane(800, 600, 2., 10.)
	assert.Nil(t, err)
	assert.Equal(t, 800, projectionPlane.Width)
	assert.Equal(t, 600, projectionPlane.Height)
	assert.Equal(t, 2., projectionPlane.PixelSize)
	assert.Equal(t, 10., projectionPlane.Distance)
	assert.Equal(t, 1., projectionPlane.Gamma)
	assert.Equal(t, false, projectionPlane.ClampOutOfGamut)
	assert.Equal(t, *math.NewColor(0., 0., 0.), projectionPlane.ClampColor)
}

func Test_NewProjectionPlane_WhenRequiredParamsAreMissing(t *testing.T) {
	projectionPlane, err := NewProjectionPlane(800)
	assert.Equal(t, fmt.Errorf("Invalid number of arguments: %d arguments received", 1), err)
	assert.Equal(t, 0, projectionPlane.Width)
	assert.Equal(t, 0, projectionPlane.Height)
	assert.Equal(t, 1., projectionPlane.PixelSize)
	assert.Equal(t, 500., projectionPlane.Distance)
	assert.Equal(t, 1., projectionPlane.Gamma)
	assert.Equal(t, false, projectionPlane.ClampOutOfGamut)
	assert.Equal(t, *math.NewColor(0., 0., 0.), projectionPlane.ClampColor)
}

func Test_NewProjectionPlane_WhenNoParamsAreReceived(t *testing.T) {
	projectionPlane, err := NewProjectionPlane()
	assert.Equal(t, fmt.Errorf("Invalid number of arguments: %d arguments received", 0), err)
	assert.Equal(t, 0, projectionPlane.Width)
	assert.Equal(t, 0, projectionPlane.Height)
	assert.Equal(t, 1., projectionPlane.PixelSize)
	assert.Equal(t, 500., projectionPlane.Distance)
	assert.Equal(t, 1., projectionPlane.Gamma)
	assert.Equal(t, false, projectionPlane.ClampOutOfGamut)
	assert.Equal(t, *math.NewColor(0., 0., 0.), projectionPlane.ClampColor)
}

func Test_NewProjectionPlane_WhenListOfParamsIsExceeded(t *testing.T) {
	projectionPlane, err := NewProjectionPlane(800, 600, 2., 10., 10.)
	assert.Equal(t, fmt.Errorf("Invalid number of arguments: %d arguments received", 5), err)
	assert.Equal(t, 0, projectionPlane.Width)
	assert.Equal(t, 0, projectionPlane.Height)
	assert.Equal(t, 1., projectionPlane.PixelSize)
	assert.Equal(t, 500., projectionPlane.Distance)
	assert.Equal(t, 1., projectionPlane.Gamma)
	assert.Equal(t, false, projectionPlane.ClampOutOfGamut)
	assert.Equal(t, *math.NewColor(0., 0., 0.), projectionPlane.ClampColor)
}

func Test_ProjectionPlane_SetPixelWhenColorOverflowsAndOutOfGamutIsSet(t *testing.T) {
	projectionPlane, err := NewProjectionPlane(200, 200, 1., 10.)
	assert.Nil(t, err)

	projectionPlane.ClampOutOfGamut = true

	projectionPlane.SetPixel(5, 5, *math.NewColor(2., .5, .5))
	assert.Equal(t, *math.NewColor(0., 0., 0.), projectionPlane.Pixels[5][5])
}

func Test_ProjectionPlane_SetPixelWhenColorOverflowsAndOutOfGamutIsNotSet(t *testing.T) {
	projectionPlane, err := NewProjectionPlane(200, 200, 1., 10.)
	assert.Nil(t, err)

	projectionPlane.SetPixel(5, 5, *math.NewColor(2., .5, .5))
	assert.Equal(t, *math.NewColor(1., .25, .25), projectionPlane.Pixels[5][5])
}

func Test_ProjectionPlane_SetPixelWhenColorDoesNotOverflow(t *testing.T) {
	projectionPlane, err := NewProjectionPlane(200, 200, 1., 10.)
	assert.Nil(t, err)

	projectionPlane.SetPixel(5, 5, *math.NewColor(1., .5, .5))
	assert.Equal(t, *math.NewColor(1., .5, .5), projectionPlane.Pixels[5][5])
}

func Test_ProjectionPlane_SetPixelWhenGammaIsDifferentFromOne(t *testing.T) {
	projectionPlane, err := NewProjectionPlane(200, 200, 1., 10.)
	assert.Nil(t, err)

	projectionPlane.Gamma = .5

	projectionPlane.SetPixel(5, 5, *math.NewColor(1., .5, .5))
	assert.Equal(t, *math.NewColor(1., .25, .25), projectionPlane.Pixels[5][5])
}