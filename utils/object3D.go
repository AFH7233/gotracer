package utils

const ERROR = 0.00001

type Object3D interface {
	Intersect(ray Ray) (Ray, float64, bool)
	Transform(transformMatrix Transformation)
}

type VisibleObject struct {
	Geometry Object3D
	Material Material
}
