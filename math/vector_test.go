package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewVector_WhenANewVectorIsCreated(t *testing.T) {
	v := NewVector(1., 2., 3.)
	assert.Equal(t, 1., v.x)
	assert.Equal(t, 2., v.y)
	assert.Equal(t, 3., v.z)
}

func Test_Add_WhenSecondOperandHasOnlyPositiveCoordinates(t *testing.T) {
	v := NewVector(1., 1., 1.)
	w := NewVector(2., 3., 1.)

	result := v.Add(w)

	assert.Equal(t, 3., result.x)
	assert.Equal(t, 4., result.y)
	assert.Equal(t, 2., result.z)
}

func Test_Add_WhenSecondOperandHasNegativeCoordinates(t *testing.T) {
	v := NewVector(1., 1., 1.)
	w := NewVector(2., -3., -1.)

	result := v.Add(w)

	assert.Equal(t, 3., result.x)
	assert.Equal(t, -2., result.y)
	assert.Equal(t, 0., result.z)
}

func Test_Sub_WhenSecondOperandHasOnlyPositiveCoordinates(t *testing.T) {
	v := NewVector(1., 1., 1.)
	w := NewVector(2., 3., 1.)

	result := v.Sub(w)

	assert.Equal(t, -1., result.x)
	assert.Equal(t, -2., result.y)
	assert.Equal(t, 0., result.z)
}

func Test_Sub_WhenSecondOperandHasNegativeCoordinates(t *testing.T) {
	v := NewVector(1., 1., 1.)
	w := NewVector(2., -3., -1.)

	result := v.Sub(w)

	assert.Equal(t, -1., result.x)
	assert.Equal(t, 4., result.y)
	assert.Equal(t, 2., result.z)
}

func Test_Multiply_WhenSecondOperandIsPositive(t *testing.T) {
	v := NewVector(1., 2., 3.)
	result := v.Multiply(5.)

	assert.Equal(t, 5., result.x)
	assert.Equal(t, 10., result.y)
	assert.Equal(t, 15., result.z)
}

func Test_Multiply_WhenSecondOperandIsNegative(t *testing.T) {
	v := NewVector(1., 2., 3.)
	result := v.Multiply(-1.)

	assert.Equal(t, -1., result.x)
	assert.Equal(t, -2., result.y)
	assert.Equal(t, -3., result.z)
}

func Test_Dot(t *testing.T) {
	v := NewVector(1., 1., 1.)
	w := NewVector(2., -3., -1.)

	result := v.Dot(w)

	assert.Equal(t, -2., result)
}
