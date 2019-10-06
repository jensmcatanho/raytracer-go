package sampler

import (
	"errors"
	"fmt"
	"jensmcatanho/raytracer-go/math"
	"math/rand"
)

type Sampler struct {
	numSamples int
	numSets    int

	samples         []math.Vector
	shuffledIndices []int

	count int
	jump  int
}

func NewSampler(args ...interface{}) (*Sampler, error) {
	numSamples, numSets, err := samplerParams(args)

	sampler := &Sampler{
		numSamples:      numSamples,
		numSets:         numSets,
		samples:         make([]math.Vector, numSamples*numSets),
		shuffledIndices: rand.Perm(numSamples),
		count:           0,
		jump:            0,
	}

	return sampler, err
}

func samplerParams(args []interface{}) (numSamples, numSets int, err error) {
	numSamples = 1
	numSets = 1

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

			numSamples = param

		case 1:
			param, ok := paramInterface.(int)
			if !ok {
				err = errors.New("2nd parameter is not of type int")
				return
			}

			numSets = param
		}
	}
	return
}
