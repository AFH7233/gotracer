package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"math/rand"
	"os"
	"time"

	"github.com/AFH7233/gotracer/utils"
)

var RAYS_PER_PIXEL = 5000
var BOUNCES = 5

func main() {
	fmt.Println("Starting render")
	start := time.Now()
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

	world := utils.NewSphere(utils.NewVector(0.0, -20.0, -10.0), 20.0)
	worldMaterial := utils.Material{Color: color.RGBA{255, 100, 100, 255}, Emitance: utils.NewNormal(0.0, 0.0, 0.0), PScatter: 1.0, Nt: 0.0001, ProbReflected: 0.0}

	sphere := utils.NewSphere(utils.NewVector(0.0, 5.5, -10.0), 4.0)
	simpleMaterial := utils.Material{Color: color.RGBA{255, 255, 255, 255}, Emitance: utils.NewNormal(0.0, 0.0, 0.0), PScatter: 0.1, Nt: 1.5, ProbReflected: 0.0}

	light := utils.NewSphere(utils.NewVector(10.0, 30, -10.0), 5.0)
	brightMaterial := utils.Material{Color: color.RGBA{255, 255, 255, 255}, Emitance: utils.NewNormal(10.0, 10.0, 10.0), PScatter: 0.5, Nt: 0.0001, ProbReflected: 0.0}

	visibleObjects := []utils.VisibleObject{
		{Geometry: &sphere, Material: simpleMaterial},
		{Geometry: &light, Material: brightMaterial},
		{Geometry: &world, Material: worldMaterial},
	}

	for _, object := range visibleObjects {
		object.Geometry.Transform(lookAt)
	}

	nc := []float64{1.0}
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			accumulatorColor := utils.NewVector(0.0, 0.0, 0.0)
			for k := 0; k < RAYS_PER_PIXEL; k++ {
				r1 := 2.0 * (random.Float64())
				r2 := 2.0 * (random.Float64())
				randX := randomizeRay(r1)
				randY := randomizeRay(r2)
				x := (2.0 * (float64(i) + randX) / float64(height)) - aspect
				y := -((2.0 * (float64(j) + randY) / float64(height)) - 1.0)
				origin := utils.NewVector(0.0, 0.0, 0.0)
				direction := utils.NewNormal(-x, y, d)
				ray := utils.NewRay(origin, direction)
				rayColor := renderColor(visibleObjects, ray, 0, nc)
				accumulatorColor = accumulatorColor.Add(rayColor)
			}

			pixelColor := accumulatorColor.Scale(1.0 / float64(RAYS_PER_PIXEL))
			img.Set(i, j, utils.Vector2Color(pixelColor))
		}
	}

	f, _ := os.Create("image.png")
	png.Encode(f, img)
	duration := time.Since(start)
	fmt.Println(duration)
}

func renderColor(objects []utils.VisibleObject, ray utils.Ray, bounces int, nco []float64) utils.Vector {
	minDistance := math.Inf(1)
	hittedObject := -1
	var distance float64 = 0.0
	var resultRay utils.Ray
	var reflectedRay utils.Ray
	for i, object := range objects {
		var isIntersected bool
		reflectedRay, distance, isIntersected = object.Geometry.Intersect(ray)
		if isIntersected {
			if distance < minDistance {
				minDistance = distance
				hittedObject = i
				resultRay = reflectedRay
			}
		}
	}

	if hittedObject != -1 {
		objectColor := utils.Color2Vector(objects[hittedObject].Material.Color).Scale(0.99)
		objectEmitance := objects[hittedObject].Material.Emitance
		surfaceNormal := resultRay.GetDirection().Scale(-1)
		normalDirection := ray.GetDirection().Dot(surfaceNormal)

		var correctedNormal utils.Vector
		var isInside bool
		if normalDirection < 0.0 {
			isInside = true
			correctedNormal = surfaceNormal
		} else {
			isInside = false
			correctedNormal = resultRay.GetDirection()
		}

		surfaceRay := utils.NewRay(resultRay.GetOrigin(), correctedNormal)
		specularRay := ray.SpecularReflection(surfaceRay).GetDirection()
		diffuseRay := ray.DiffuseReflection(surfaceRay).GetDirection()
		pScatter := objects[hittedObject].Material.PScatter
		pSpecular := 1.0 - pScatter
		var reflectedVector utils.Vector

		nt := objects[hittedObject].Material.Nt

		var nnt float64
		if isInside {
			nc := nco[len(nco)-2]
			nnt = nt / nc
		} else {
			nc := nco[len(nco)-1]
			nnt = nc / nt
		}
		ddn := ray.GetDirection().Dot(correctedNormal)
		cos2t := 1 - (nnt * nnt * (1 - (ddn * ddn)))

		enteredObject := false
		if cos2t < 0.0 {
			reflectedVector = specularRay.Scale(pSpecular).Add(diffuseRay.Scale(pScatter)).Normalize()
		} else {
			refractedRay := ray.GetDirection().Sub(correctedNormal.Scale(ddn)).Scale(nnt).Sub(correctedNormal.Scale(math.Sqrt(cos2t)))
			probReflected := rand.Float64()
			if probReflected < objects[hittedObject].Material.ProbReflected && !isInside {
				reflectedVector = specularRay.Scale(pSpecular).Add(diffuseRay.Scale(pScatter)).Normalize()
			} else {
				enteredObject = true
				reflectedVector = refractedRay
			}
		}

		if bounces > BOUNCES {
			return objectEmitance
		}
		reflectedRay := utils.NewRay(resultRay.GetOrigin(), reflectedVector)
		bounces += 1
		var recursionColor utils.Vector
		if isInside {
			if enteredObject {
				nco = nco[:len(nco)-1]
			}
		} else {
			if enteredObject {
				nco = append(nco, nt)
			}
		}
		recursionColor = renderColor(objects, reflectedRay, bounces, nco)
		return recursionColor.Multiply(objectColor).Add(objectEmitance)
	}
	return utils.NewNormal(0.01, 0.05, 0.08)
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
