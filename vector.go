package main

import (
	"math"

	"github.com/shockdev04/altv-go-pkg/alt"
)

func GetForwardVector(rot alt.Rotation) alt.Position {
	z := float64(-rot.Z)
	x := float64(rot.X)
	num := math.Abs(math.Cos(x))
	return alt.Position{
		X: float32(-math.Sin(z) * num),
		Y: float32(math.Cos(z) * num),
		Z: float32(math.Sin(x)),
	}
}

func GetPositionInFront(pos alt.Position, rot alt.Rotation, dist float32) alt.Position {
	fwd := GetForwardVector(rot)
	return alt.Position{
		X: pos.X + fwd.X*dist,
		Y: pos.Y + fwd.Y*dist,
		Z: pos.Z,
	}
}
