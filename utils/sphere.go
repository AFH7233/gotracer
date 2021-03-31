package utils

import "math"

type Sphere struct {
	origin Vector
	radius float64
}

func NewSphere(origin Vector, radius float64) (sphere Sphere) {
	sphere.origin = origin
	sphere.radius = radius
	return
}

func (sphere *Sphere) Intersect(ray Ray) (result Ray, distance float64, isHitted bool) {
	raySphere := ray.origin.Sub(sphere.origin)
	b := ray.direction.Dot(raySphere)
	discriminante := b*b - (raySphere.Dot(raySphere) - sphere.radius*sphere.radius)
	if discriminante < ERROR {
		isHitted = false
		return
	} else {
		c := math.Sqrt(discriminante)

		if (-b - c) > ERROR {
			distance = (-b - c)
		} else if (-b + c) > ERROR {
			distance = (-b + c)
		} else {
			isHitted = false
			return
		}
		isHitted = true
		point := ray.GetRayPoint(distance)
		result.origin = point
		result.direction = point.Sub(sphere.origin).Normalize()
		return
	}
}

func (sphere *Sphere) Transform(transformMatrix Transformation) {
	newOrigin := sphere.origin.Transform(transformMatrix)
	sphere.origin = newOrigin
}
