package main

import (
	"image"
	"image/png"
	"math"
	"math/rand"
	"os"
	"github.com/AFH7233/gotracer/utils"
)

var RAYS_PER_PIXEL = 100

func main() {
	width := 640
	height := 480
	random := rand.New(rand.NewSource(99))
	aspect := float64(width) / float64(height)

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	camera := utils.NewCamera(
		utils.NewNormal(0.0, 1.0, 0.0),
		utils.NewVector(0.0, 5.0, -40.0),
		45.0,
	)
	lookAt := camera.GetLookAt(utils.NewVector(0.0, 5.0, 20.0))
	d := camera.GetDistanceFromScreen(aspect)

	sphere := utils.NewSphere(utils.NewVector(0.0, 7.5, -10.0), 2.0)
	objects := []utils.Object3D{&sphere}

	for _, object := range objects {
		object.Transform(lookAt)
	}

	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			accumulatorColor := utils.NewVector(0.0, 0.0, 0.0)
			for k := 0; k < RAYS_PER_PIXEL ; k++ {
				r1 := 2.0 * (random.Float64())
				r2 := 2.0 * (random.Float64())
				randX := randomizeRay(r1)
				randY := randomizeRay(r2)
				x := (2.0 * (float64(i) + randX) / float64(height)) - aspect
				y := -((2.0 * (float64(j) + randY) / float64(height)) - 1.0)
				origin := utils.NewVector(0.0, 0.0, 0.0)
				direction := utils.NewNormal(x, y, -d)
				ray := utils.NewRay(origin, direction)
				rayColor := renderColor(objects, ray)
				accumulatorColor = accumulatorColor.Add(rayColor)
			}
			pixelColor := accumulatorColor.Scale(1.0/float64(RAYS_PER_PIXEL))
			img.Set(i, j, utils.Vector2Color(pixelColor))
		}
	}

	f, _ := os.Create("image.png")
	png.Encode(f, img)
}

func renderColor(objects []utils.Object3D, ray utils.Ray) utils.Vector {
	minDistance := math.Inf(1)
	hittedObject := -1
	var distance float64 = 0.0
	var resultRay utils.Ray
	for i, object := range objects {
		var isIntersected bool
		resultRay, distance, isIntersected = object.Intersect(ray)
		if isIntersected {
			if distance < minDistance {
				minDistance = distance
				hittedObject = i
			}
		}
	}

	if hittedObject != -1 {
		return resultRay.GetDirection()

	}
	return utils.NewVector(0.0, 0.0, 0.0)
}

func randomizeRay(randomNumber float64) float64 {
	var deviation float64
	if randomNumber < 1.0 {
		deviation = math.Sqrt(randomNumber) - 1.0
	} else {
		deviation = 1.0 - math.Sqrt(2.0-randomNumber)
	}
	return deviation
}
