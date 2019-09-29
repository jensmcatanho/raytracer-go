package main

import (
	"fmt"
	"jensmcatanho/raytracer-go/math"
)

func main() {
	origin := math.NewVector(0., 0., 0.)
	fmt.Println(fmt.Sprintf("Raytracer Go!\nOrigin: %+v", origin))
}
