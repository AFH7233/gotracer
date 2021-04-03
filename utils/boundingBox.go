package utils

import (
	"math"
)

type Box struct {
	MinX float64
	MaxX float64
	MinY float64
	MaxY float64
	MinZ float64
	MaxZ float64
}

func (a Box) Collide(b Box) bool {
	if b.MinX > a.MaxX || b.MaxX < a.MinX || b.MinY > a.MaxY || b.MaxY < a.MinY || b.MinZ > a.MaxZ || b.MaxZ < a.MinZ {
		return false
	}
	return true
}

func (box Box) Intersect(ray Ray) bool {
	var dirfrac Vector
	// r.dir is unit direction vector of ray
	dirfrac.x = 1.0 / ray.direction.x
	dirfrac.y = 1.0 / ray.direction.y
	dirfrac.z = 1.0 / ray.direction.z
	// lb is the corner of AABB with minimal coordinates - left bottom, rt is maximal corner
	// ray.origin is origin of ray
	t1 := (box.MinX - ray.origin.x) * dirfrac.x
	t2 := (box.MaxX - ray.origin.x) * dirfrac.x
	t3 := (box.MinY - ray.origin.y) * dirfrac.y
	t4 := (box.MaxY - ray.origin.y) * dirfrac.y
	t5 := (box.MinZ - ray.origin.z) * dirfrac.z
	t6 := (box.MaxZ - ray.origin.z) * dirfrac.z

	tmin := math.Max(math.Max(math.Min(t1, t2), math.Min(t3, t4)), math.Min(t5, t6))
	tmax := math.Min(math.Min(math.Max(t1, t2), math.Max(t3, t4)), math.Max(t5, t6))

	// if tmax < 0, ray (line) is intersecting AABB, but the whole AABB is behind us
	if tmax < 0 {
		return false
	}

	// if tmin > tmax, ray doesn't intersect AABB
	if tmin > tmax {
		return false
	}

	return true
}
