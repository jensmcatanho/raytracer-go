package scene

import (
	"fmt"
	"jensmcatanho/raytracer-go/math"
	"jensmcatanho/raytracer-go/object"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sceneImpl_Initialize_WhenMinimumParamsArePassed(t *testing.T) {
	err := GetInstance().Initialize()
	assert.Nil(t, err)
	assert.Empty(t, GetInstance().ObjectList())
}

func Test_sceneImpl_Initialize_WhenOnlyBackgroundColorIsPassed(t *testing.T) {
	err := GetInstance().Initialize(*math.NewColor(0., 0., 0.))
	assert.Nil(t, err)
	assert.Equal(t, 0, len(GetInstance().ObjectList()))
}

func Test_sceneImpl_Initialize_WhenAllParamsArePassed(t *testing.T) {
	err := GetInstance().Initialize(*math.NewColor(0., 0., 0.), 10)
	assert.Nil(t, err)
	assert.Equal(t, 10, len(GetInstance().ObjectList()))
}

func Test_sceneImpl_Initialize_WhenListOfParamsIsExceeded(t *testing.T) {
	err := GetInstance().Initialize(*math.NewColor(0., 0., 0.), 10, 10)
	assert.Equal(t, fmt.Errorf("Invalid number of arguments: %d arguments received", 3), err)
}

func Test_sceneImpl_Initialize_WhenFirstParamHasInvalidType(t *testing.T) {
	err := GetInstance().Initialize(10)
	assert.Equal(t, fmt.Errorf("1st parameter is not of type math.Color"), err)
}

func Test_sceneImpl_Initialize_WhenSecondParamHasInvalidType(t *testing.T) {
	err := GetInstance().Initialize(*math.NewColor(0., 0., 0.), 10.)
	assert.Equal(t, fmt.Errorf("2nd parameter is not of type int"), err)
}

func Test_sceneImpl_AddObject_WhenObjectListSizeIsNotSpecified(t *testing.T) {
	err := GetInstance().Initialize()
	assert.Nil(t, err)

	sphereCenter := math.NewVector(0., 0., 10.)
	sphere := object.NewSphere(*sphereCenter, 2)
	GetInstance().AddObject(sphere)
	assert.NotEmpty(t, GetInstance().ObjectList())
}

func Test_sceneImpl_AddObject_WhenObjectListSizeIsSpecified(t *testing.T) {
	err := GetInstance().Initialize(*math.NewColor(0., 0., 0.), 10)
	assert.Nil(t, err)

	sphereCenter := math.NewVector(0., 0., 10.)
	sphere := object.NewSphere(*sphereCenter, 2)
	GetInstance().AddObject(sphere)
	assert.NotEmpty(t, GetInstance().ObjectList())
}
