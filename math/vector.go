package math

// Vector is an abstraction of a 3D vector
type Vector struct {
	x, y, z float64
}

// NewVector creates a Vector structure
func NewVector(x, y, z float64) *Vector {
	return &Vector{
		x: x,
		y: y,
		z: z,
	}
}

// Add adds two vectors and returns a new Vector with the result
func (v *Vector) Add(w *Vector) *Vector {
	return &Vector{
		x: v.x + w.x,
		y: v.y + w.y,
		z: v.z + w.z,
	}
}

// Sub subtracts two vectors and returns a new Vector with the result
func (v *Vector) Sub(w *Vector) *Vector {
	return &Vector{
		x: v.x - w.x,
		y: v.y - w.y,
		z: v.z - w.z,
	}
}

// Multiply multiplies a Vector by a constant and returns a new Vector with the result
func (v *Vector) Multiply(k float64) *Vector {
	return &Vector{
		x: v.x * k,
		y: v.y * k,
		z: v.z * k,
	}
}

// Dot calculates the dot product between two Vectors
func (v *Vector) Dot(w *Vector) float64 {
	return v.x*w.x + v.y*w.y + v.z*w.z
}
