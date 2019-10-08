package sampler

import (
	"errors"
	"fmt"
	"jensmcatanho/raytracer-go/math"
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

func shuffleX(numSamples, numSets int, samples *[]math.Vector) {
	for i := 0; i < numSets; i++ {
		for j := 0; i < numSamples-1; j++ {
			index := rand.Int()%numSamples + i*numSamples
			value := (*samples)[j+i*numSamples+1].X
			(*samples)[j+i*numSamples+1].X = (*samples)[index].X
			(*samples)[index].X = value
		}
	}
}

func shuffleY(numSamples, numSets int, samples *[]math.Vector) {
	for i := 0; i < numSets; i++ {
		for j := 0; i < numSamples-1; j++ {
			index := rand.Int()%numSamples + i*numSamples
			value := (*samples)[j+i*numSamples+1].Y
			(*samples)[j+i*numSamples+1].Y = (*samples)[index].Y
			(*samples)[index].Y = value
		}
	}
}
