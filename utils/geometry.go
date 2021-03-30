package utils

type Object3D interface {
	Intersect(ray Ray) (Ray, float64, bool)
	Transform(transformMatrix Transformation)
}
