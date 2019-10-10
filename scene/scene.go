package scene

import (
	"errors"
	"fmt"
	"jensmcatanho/raytracer-go/math/color"
	"jensmcatanho/raytracer-go/math/geometry"
	"sync"
)

var (
	once     sync.Once
	instance Scene
)

type Scene interface {
	Initialize(args ...interface{}) error
	AddObject(obj geometry.Renderable)
	ObjectList() []geometry.Renderable
}

type sceneImpl struct {
	backgroundColor color.Color
	objects         []geometry.Renderable
}

func GetInstance() Scene {
	once.Do(func() {
		instance = &sceneImpl{}
	})

	return instance
}

func (s *sceneImpl) Initialize(args ...interface{}) error {
	backgroundColor, numObjects, err := sceneParams(args)
	objects := make([]geometry.Renderable, numObjects)

	s.backgroundColor = backgroundColor
	s.objects = objects

	return err
}

func sceneParams(args []interface{}) (backgroundColor color.Color, numObjects int, err error) {
	backgroundColor = *color.NewColor(0., 0., 0.)
	numObjects = 0

	if len(args) > 2 {
		err = fmt.Errorf("Invalid number of arguments: %d arguments received", len(args))
		return
	}

	for i, paramInterface := range args {
		switch i {
		case 0:
			param, ok := paramInterface.(color.Color)
			if !ok {
				err = errors.New("1st parameter is not of type color.Color")
				return
			}

			backgroundColor = param

		case 1:
			param, ok := paramInterface.(int)
			if !ok {
				err = errors.New("2nd parameter is not of type int")
				return
			}

			numObjects = param
		}
	}
	return
}

func (s *sceneImpl) AddObject(obj geometry.Renderable) {
	s.objects = append(s.objects, obj)
}

func (s *sceneImpl) ObjectList() []geometry.Renderable {
	return s.objects
}
