package utils

import (
	"image/color"
)

type Material struct {
	Color    color.RGBA
	Emitance Vector
	PScatter float64
}
