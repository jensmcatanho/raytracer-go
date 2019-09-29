package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewColor_WhenANewColorIsCreated(t *testing.T) {
	color := NewColor(1., .5, .3)
	assert.Equal(t, float32(1.), color.r)
	assert.Equal(t, float32(.5), color.g)
	assert.Equal(t, float32(.3), color.b)
}
