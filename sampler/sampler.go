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
	Samples, Sets, err := samplerParams(args)

	sampler := &Sampler{
		Samples:         Samples,
		Sets:            Sets,
		method:          Regular,
		shuffledIndices: rand.Perm(Samples),
		count:           0,
		jump:            0,
	}

	return sampler, err
}

func samplerParams(args []interface{}) (Samples, Sets int, err error) {
	Samples = 1
	Sets = 1

	if len(args) > 2 {
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

			Samples = param

		case 1:
			param, ok := paramInterface.(int)
			if !ok {
				err = errors.New("2nd parameter is not of type int")
				return
			}

			Sets = param
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
	s.count = s.count + 1

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

	fmt.Println(samples)
}
