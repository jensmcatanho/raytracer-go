package scene

import (
	"errors"
	"fmt"
	"jensmcatanho/raytracer-go/math"
	"jensmcatanho/raytracer-go/object"
	"sync"
)

var (
	once     sync.Once
	instance Scene
)

type Scene interface {
	Initialize(args ...interface{}) error
	AddObject(obj object.Renderable)
	ObjectList() []object.Renderable
}

type sceneImpl struct {
	backgroundColor math.Color
	objects         []object.Renderable
}

func GetInstance() Scene {
	once.Do(func() {
		instance = &sceneImpl{}
	})

	return instance
}

func (s *sceneImpl) Initialize(args ...interface{}) error {
	backgroundColor, numObjects, err := sceneParams(args)
	objects := make([]object.Renderable, numObjects)

	s.backgroundColor = backgroundColor
	s.objects = objects

	return err
}

func sceneParams(args []interface{}) (backgroundColor math.Color, numObjects int, err error) {
	backgroundColor = *math.NewColor(0., 0., 0.)
	numObjects = 0

	if len(args) > 2 {
		err = fmt.Errorf("Invalid number of arguments: %d arguments received", len(args))
		return
	}

	for i, paramInterface := range args {
		switch i {
		case 0:
			param, ok := paramInterface.(math.Color)
			if !ok {
				err = errors.New("1st parameter is not of type math.Color")
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

func (s *sceneImpl) AddObject(obj object.Renderable) {
	s.objects = append(s.objects, obj)
}

func (s *sceneImpl) ObjectList() []object.Renderable {
	return s.objects
}
