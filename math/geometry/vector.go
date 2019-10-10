package geometry

import (
	"math"
)

// Vector is a structure that represents a 3D vector
type Vector struct {
	X, Y, Z float64
}

// NewVector creates a Vector structure
func NewVector(x, y, z float64) *Vector {
	return &Vector{
		X: x,
		Y: y,
		Z: z,
	}
}

// Add adds two vectors and returns a new Vector with the result
func (v *Vector) Add(w *Vector) *Vector {
	return &Vector{
		X: v.X + w.X,
		Y: v.Y + w.Y,
		Z: v.Z + w.Z,
	}
}

// Sub subtracts two vectors and returns a new Vector with the result
func (v *Vector) Sub(w *Vector) *Vector {
	return &Vector{
		X: v.X - w.X,
		Y: v.Y - w.Y,
		Z: v.Z - w.Z,
	}
}

// Multiply multiplies a Vector by a constant and returns a new Vector with the result
func (v *Vector) Multiply(k float64) *Vector {
	return &Vector{
		X: v.X * k,
		Y: v.Y * k,
		Z: v.Z * k,
	}
}

// Dot calculates the dot product between two Vectors
func (v *Vector) Dot(w *Vector) float64 {
	return v.X*w.X + v.Y*w.Y + v.Z*w.Z
}

// Length returns the length of the Vector
func (v *Vector) Length() float64 {
	return math.Sqrt(math.Pow(v.X, 2) + math.Pow(v.Y, 2) + math.Pow(v.Z, 2))
}

// Normalize normalizes the Vector
func (v *Vector) Normalize() {
	norm := v.Length()

	v.X = v.X / norm
	v.Y = v.Y / norm
	v.Z = v.Z / norm
}

// Cross calculates the cross product between two Vectors
func (v *Vector) Cross(w *Vector) *Vector {
	crossProduct := NewVector(0., 0., 0.)

	crossProduct.X = v.Y*w.Z - v.Z*w.Y
	crossProduct.Y = v.X*w.Z - v.Z*w.X
	crossProduct.Z = v.X*w.Y - v.Y*w.X

	return crossProduct
}
