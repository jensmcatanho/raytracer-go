package sampler

import (
	"jensmcatanho/raytracer-go/math"
	stdMath "math"
	"math/rand"
	"time"
)

func Regular(numSamples, numSets int, samples *[]math.Vector) {
	n := int(stdMath.Sqrt(float64(numSamples)))

	for i := 0; i < numSets; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				x := (float64(k) + 0.5) / float64(n)
				y := (float64(j) + 0.5) / float64(n)

				*samples = append(*samples, *math.NewVector(float64(x), float64(y), 0.))
			}
		}
	}
}

func Random(numSamples, numSets int, samples *[]math.Vector) {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < numSets; i++ {
		for j := 0; j < numSamples; j++ {
			*samples = append(*samples, *math.NewVector(rand.Float64(), rand.Float64(), .0))
		}
	}
}

// Hammersley generates sample using Hammersley points
// http://holger.dammertz.org/stuff/notes_HammersleyOnHemisphere.html
func Hammersley(numSamples, numSets int, samples *[]math.Vector) {
	phi := func(j int) float64 {
		x := 0.
		f := .5

		for j > 0 {
			x += f * float64(j%2)
			j /= 2
			f *= .5
		}

		return x
	}

	for i := 0; i < numSets; i++ {
		for j := 0; j < numSamples; j++ {
			x := float64(j) / float64(numSamples)
			y := phi(j)

			*samples = append(*samples, *math.NewVector(x, y, 0.))
		}
	}
}

func NRooks(numSamples, numSets int, samples *[]math.Vector) {
	for i := 0; i < numSets; i++ {
		for j := 0; j < numSamples; j++ {
			x := (float64(j) + rand.Float64()) / float64(numSamples)
			y := (float64(j) + rand.Float64()) / float64(numSamples)

			*samples = append(*samples, *math.NewVector(x, y, .0))
		}
	}

	shuffleX(numSamples, numSets, samples)
	shuffleY(numSamples, numSets, samples)
}

func Jittered(numSamples, numSets int, samples *[]math.Vector) {
	n := stdMath.Sqrt(float64(numSamples))

	for i := 0; i < numSets; i++ {
		for j := 0; float64(j) < n; j++ {
			for k := 0; float64(k) < n; k++ {
				x := (float64(k) + rand.Float64()) / n
				y := (float64(j) + rand.Float64()) / n

				*samples = append(*samples, *math.NewVector(x, y, .0))
			}
		}
	}
}
