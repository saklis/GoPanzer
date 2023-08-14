package Structs

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func RotateByAngle(vector rl.Vector2, angle float32) rl.Vector2 {
	radians := float32(math.Pi * angle / 180.0)
	cosTheta := float32(math.Cos(float64(radians)))
	sinTheta := float32(math.Sin(float64(radians)))
	rotatedX := vector.X*cosTheta - vector.Y*sinTheta
	rotatedY := vector.X*sinTheta + vector.Y*cosTheta
	vector.X = rotatedX
	vector.Y = rotatedY
	return vector
}
