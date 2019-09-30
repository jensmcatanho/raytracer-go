package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewRay_WhenANewRayIsCreated(t *testing.T) {
	origin := NewVector(0., 0., 0.)
	direction := NewVector(0., 0., 1.)

	ray := NewRay(*origin, *direction)

	assert.Equal(t, *origin, ray.Origin)
	assert.Equal(t, *direction, ray.Direction)
}
