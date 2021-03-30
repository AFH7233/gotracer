package utils

import (
	"math"
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

	side := camera.up.Cross(forward)
	side = side.Normalize()

	up := forward.Cross(side)
	up = up.Normalize()

	ejes := Transformation{
		{side.x, side.y, side.z, side.q},
		{up.x, up.y, up.z, up.q},
		{forward.x, forward.y, forward.z, forward.q},
		{0.0, 0.0, 0.0, 1.0},
	}

	translacion := Transformation{
		{1.0, 0.0, 0.0, -camera.origin.x},
		{0.0, 1.0, 0.0, -camera.origin.y},
		{0.0, 0.0, 1.0, -camera.origin.z},
		{0.0, 0.0, 0.0, 1.0},
	}

	lookAt := ejes.Combine(translacion)

	return lookAt
}

func (camera Camera) GetDistanceFromScreen(aspect float64) float64 {
	return (aspect*aspect + 1) / math.Tan(math.Pi*camera.fov/180.0)
}
