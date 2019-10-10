package color

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewColor_WhenANewColorIsCreated(t *testing.T) {
	color := NewColor(1., .5, .3)
	assert.Equal(t, 1., color.r)
	assert.Equal(t, .5, color.g)
	assert.Equal(t, .3, color.b)
}

func Test_Color_Multiply_WhenSecondOperandIsPositive(t *testing.T) {
	v := NewColor(1., 2., 3.)
	result := v.Multiply(5.)

	assert.Equal(t, 5., result.r)
	assert.Equal(t, 10., result.g)
	assert.Equal(t, 15., result.b)
}

func Test_Color_Multiply_WhenSecondOperandIsNegative(t *testing.T) {
	v := NewColor(1., 2., 3.)
	result := v.Multiply(-1.)

	assert.Equal(t, -1., result.r)
	assert.Equal(t, -2., result.g)
	assert.Equal(t, -3., result.b)
}
