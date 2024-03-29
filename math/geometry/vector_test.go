package geometry

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewVector_WhenANewVectorIsCreated(t *testing.T) {
	v := NewVector(1., 2., 3.)
	assert.Equal(t, 1., v.X)
	assert.Equal(t, 2., v.Y)
	assert.Equal(t, 3., v.Z)
}

func Test_Add_WhenSecondOperandHasOnlyPositiveCoordinates(t *testing.T) {
	v := NewVector(1., 1., 1.)
	w := NewVector(2., 3., 1.)

	result := v.Add(w)

	assert.Equal(t, 3., result.X)
	assert.Equal(t, 4., result.Y)
	assert.Equal(t, 2., result.Z)
}

func Test_Add_WhenSecondOperandHasNegativeCoordinates(t *testing.T) {
	v := NewVector(1., 1., 1.)
	w := NewVector(2., -3., -1.)

	result := v.Add(w)

	assert.Equal(t, 3., result.X)
	assert.Equal(t, -2., result.Y)
	assert.Equal(t, 0., result.Z)
}

func Test_Sub_WhenSecondOperandHasOnlyPositiveCoordinates(t *testing.T) {
	v := NewVector(1., 1., 1.)
	w := NewVector(2., 3., 1.)

	result := v.Sub(w)

	assert.Equal(t, -1., result.X)
	assert.Equal(t, -2., result.Y)
	assert.Equal(t, 0., result.Z)
}

func Test_Sub_WhenSecondOperandHasNegativeCoordinates(t *testing.T) {
	v := NewVector(1., 1., 1.)
	w := NewVector(2., -3., -1.)

	result := v.Sub(w)

	assert.Equal(t, -1., result.X)
	assert.Equal(t, 4., result.Y)
	assert.Equal(t, 2., result.Z)
}

func Test_Multiply_WhenSecondOperandIsPositive(t *testing.T) {
	v := NewVector(1., 2., 3.)
	result := v.Multiply(5.)

	assert.Equal(t, 5., result.X)
	assert.Equal(t, 10., result.Y)
	assert.Equal(t, 15., result.Z)
}

func Test_Multiply_WhenSecondOperandIsNegative(t *testing.T) {
	v := NewVector(1., 2., 3.)
	result := v.Multiply(-1.)

	assert.Equal(t, -1., result.X)
	assert.Equal(t, -2., result.Y)
	assert.Equal(t, -3., result.Z)
}

func Test_Dot(t *testing.T) {
	v := NewVector(1., 1., 1.)
	w := NewVector(2., -3., -1.)

	result := v.Dot(w)

	assert.Equal(t, -2., result)
}

func Test_Length(t *testing.T) {
	v := NewVector(1., 1., 1.)
	epsilon := 0.001

	length := v.Length()

	assert.LessOrEqual(t, math.Abs(float64(1.7320508075688772)-length), epsilon)
}

func Test_Normalize(t *testing.T) {
	v := NewVector(math.Sqrt(3), math.Sqrt(3), math.Sqrt(3))
	epsilon := 0.001

	v.Normalize()

	assert.LessOrEqual(t, math.Abs(math.Sqrt(3)/3-v.X), epsilon)
	assert.LessOrEqual(t, math.Abs(math.Sqrt(3)/3-v.Y), epsilon)
	assert.LessOrEqual(t, math.Abs(math.Sqrt(3)/3-v.Z), epsilon)
}

func Test_Cross(t *testing.T) {
	v := NewVector(1., 0., 0.)
	w := NewVector(0., 0., 1.)

	result := v.Cross(w)

	assert.Equal(t, 0., result.X)
	assert.Equal(t, 1., result.Y)
	assert.Equal(t, 0., result.Z)
}
