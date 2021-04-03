package utils

import (
	"math"
	"math/rand"
)

type Ray struct {
	origin    Vector
	direction Vector
}

func NewRay(origin, direction Vector) Ray {
	origin.q = 1.0
	direction.q = 0.0
	ray := Ray{origin, direction.Normalize()}
	return ray
}

func (r Ray) GetOrigin() Vector {
	return r.origin
}

func (r Ray) GetDirection() Vector {
	return r.direction
}

func (ray Ray) GetRayPoint(distance float64) Vector {
	scaledVector := ray.direction.Scale(distance)
	result := ray.origin.Add(scaledVector)
	return result
}

func (ray Ray) SpecularReflection(normal Ray) (result Ray) {
	resultDirection := normal.direction.Scale(2.0 * ray.direction.Dot(normal.direction))
	resultDirection = normal.direction.Sub(resultDirection)
	resultDirection = resultDirection.Normalize()

	result.direction = resultDirection
	result.origin = normal.origin
	result.direction.q = 0.0
	return
}

func (ray Ray) DiffuseReflection(normal Ray) (result Ray) {
	r1 := 2.0 * math.Pi * rand.Float64()
	r2 := math.Sqrt(rand.Float64())
	op1 := Vector{x: 0.0, y: 1.0, z: 0.0, q: 0.0}
	op2 := Vector{x: 1.0, y: 0.0, z: 0.0, q: 0.0}

	var op3 Vector
	if math.Abs(normal.direction.x) > 0.1 {
		op3 = op1
	} else {
		op3 = op2
	}

	u := op3.Cross(normal.direction).Normalize()
	v := normal.direction.Cross(u).Normalize()

	ru := u.Scale(math.Cos(r1) * r2)
	rv := v.Scale(math.Sin(r1) * r2)
	rw := normal.direction.Scale(math.Sqrt(1.0 - r2))

	result.direction = ru.Add(rv).Add(rw).Normalize()
	result.origin = normal.origin
	result.direction.q = 0.0
	return
}
