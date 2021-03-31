package utils

import (
	"image/color"
	"math"
)

type Vector struct {
	x float64
	y float64
	z float64
	q float64
}

func NewVector(x, y, z float64) (result Vector) {
	result.x = x
	result.y = y
	result.z = z
	result.q = 1.0
	return
}

func NewNormal(x, y, z float64) (result Vector) {
	result.x = x
	result.y = y
	result.z = z
	result.q = 0.0
	return
}

func (a Vector) Dot(b Vector) float64 {
	return a.x*b.x + a.y*b.y + a.z*b.z
}

func (a Vector) Cross(b Vector) (result Vector) {
	result.x = (a.y * b.z) - (a.z * b.y)
	result.y = (a.z * b.x) - (a.x * b.z)
	result.z = (a.x * b.y) - (a.y * b.x)
	result.q = a.q
	return
}

func (a Vector) Normalize() (result Vector) {
	abs := math.Sqrt(a.x*a.x + a.y*a.y + a.z*a.z)
	result.x = a.x / abs
	result.y = a.y / abs
	result.z = a.z / abs
	result.q = a.q
	return
}

func (a Vector) Scale(k float64) (result Vector) {
	result.x = a.x * k
	result.y = a.y * k
	result.z = a.z * k
	result.q = a.q
	return
}

func (a Vector) Multiply(b Vector) (result Vector) {
	result.x = a.x * b.x
	result.y = a.y * b.y
	result.z = a.z * b.z
	result.q = a.q * b.q
	return
}

func (a Vector) Extend(k float64) (result Vector) {
	result.x = a.x + k
	result.y = a.y + k
	result.z = a.z + k
	result.q = a.q
	return
}

func (a Vector) Add(b Vector) (result Vector) {
	result.x = a.x + b.x
	result.y = a.y + b.y
	result.z = a.z + b.z
	result.q = a.q
	return
}

func (a Vector) Sub(b Vector) (result Vector) {
	result.x = a.x - b.x
	result.y = a.y - b.y
	result.z = a.z - b.z
	result.q = a.q
	return
}

func (a Vector) Norma() (nrma float64) {
	return math.Sqrt(math.Abs(a.Dot(a)))
}

func (a Vector) Transform(transformMatrix Transformation) (result Vector) {
	result.x = transformMatrix[0][0]*a.x + transformMatrix[0][1]*a.y + transformMatrix[0][2]*a.z + transformMatrix[0][3]*a.q
	result.y = transformMatrix[1][0]*a.x + transformMatrix[1][1]*a.y + transformMatrix[1][2]*a.z + transformMatrix[1][3]*a.q
	result.z = transformMatrix[2][0]*a.x + transformMatrix[2][1]*a.y + transformMatrix[2][2]*a.z + transformMatrix[2][3]*a.q
	result.q = transformMatrix[3][0]*a.x + transformMatrix[3][1]*a.y + transformMatrix[3][2]*a.z + transformMatrix[3][3]*a.q
	return
}

func Vector2Color(a Vector) color.RGBA {
	r := clamp(a.x)
	g := clamp(a.y)
	b := clamp(a.z)
	Ir := uint8(0.5 + math.Pow(r, 1.0/2.2)*255.0)
	Ig := uint8(0.5 + math.Pow(g, 1.0/2.2)*255.0)
	Ib := uint8(0.5 + math.Pow(b, 1.0/2.2)*255.0)
	return color.RGBA{Ir, Ig, Ib, 255}
}

func Color2Vector(c color.RGBA) (result Vector) {
	result.x = float64(c.R) / 255.0
	result.y = float64(c.G) / 255.0
	result.z = float64(c.B) / 255.0
	return
}

func clamp(x float64) float64 {
	if x < 0.0 {
		return 0.0
	} else if x > 1.0 {
		return 1.0
	} else {
		return x
	}
}
