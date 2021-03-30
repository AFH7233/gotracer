package utils

import (
	"math"
	"fmt"
)

type Camera struct {
	up     Vector
	origin Vector
	fov    float64
}

func NewCamera(up Vector, origin Vector, fov float64) (camera Camera) {
	up.q = 0.0
	camera.up = up.Normalize()
	camera.origin = origin
	camera.fov = fov
	return
}

func (camera Camera) GetLookAt(target Vector) Transformation {
	forward := camera.origin.Sub(target)
	forward.q = 0.0
	forward = forward.Normalize()
	fmt.Println(forward)

	side := camera.up.Cross(forward)
	side = side.Normalize()
	fmt.Println(side)

	up := forward.Cross(side)
	up = up.Normalize()
	fmt.Println(up)


	ejes := Transformation{
		{side.x, side.y, side.z, side.q},
		{up.x, up.y, up.z, up.q},
		{forward.x, forward.y, forward.z, forward.q},
		{0.0, 0.0, 0.0, 1.0},
	}
	fmt.Println(ejes)

	translacion := Transformation{
		{1.0, 0.0, 0.0, -camera.origin.x},
		{0.0, 1.0, 0.0, -camera.origin.y},
		{0.0, 0.0, 1.0, -camera.origin.z},
		{0.0, 0.0, 0.0, 1.0},
	}
	fmt.Println(translacion)

	lookAt := translacion.Combine(ejes)
	fmt.Println(lookAt)

	return lookAt
}

func (camera Camera) GetDistanceFromScreen(aspect float64) float64 {
	return (aspect*aspect + 1) / math.Tan(math.Pi*camera.fov/180.0)
}
