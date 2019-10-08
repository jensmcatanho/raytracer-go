package sampler

import (
	"errors"
	"fmt"
	"jensmcatanho/raytracer-go/math"
	stdMath "math"
	"math/rand"
	"time"
)

type Sampler struct {
	Samples int
	Sets    int

	method          func(int, int, *[]math.Vector)
	samples         []math.Vector
	shuffledIndices []int

	count int
	jump  int
}

func NewSampler(args ...interface{}) (*Sampler, error) {
	samples, sets, method, err := samplerParams(args)

	sampler := &Sampler{
		Samples:         samples,
		Sets:            sets,
		method:          method,
		shuffledIndices: rand.Perm(samples),
		count:           0,
		jump:            0,
	}

	return sampler, err
}

func samplerParams(args []interface{}) (samples, sets int, method func(int, int, *[]math.Vector), err error) {
	samples = 1
	sets = 1

	if len(args) > 3 {
		err = fmt.Errorf("Invalid number of arguments: %d arguments received", len(args))
		return
	}

	for i, paramInterface := range args {
		switch i {
		case 0:
			param, ok := paramInterface.(int)
			if !ok {
				err = errors.New("1st parameter is not of type int")
				return
			}

			samples = param

		case 1:
			param, ok := paramInterface.(int)
			if !ok {
				err = errors.New("2nd parameter is not of type int")
				return
			}

			sets = param

		case 2:
			param, ok := paramInterface.(func(int, int, *[]math.Vector))
			if !ok {
				err = errors.New("2nd parameter is not of type int")
				return
			}

			method = param
		}
	}
	return
}

func (s *Sampler) Sample() {
	s.method(s.Samples, s.Sets, &s.samples)
}

func (s *Sampler) SampleUnitSquare() math.Vector {
	s.setJump()

	sample := s.samples[s.jump+s.shuffledIndices[s.jump+s.count%s.Samples]]
	s.count++

	return sample
}

func (s *Sampler) setJump() {
	rand.Seed(time.Now().UnixNano())

	if s.count%s.Samples == 0 {
		s.jump = (rand.Int() % s.Sets) * s.Samples
	}
}

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
