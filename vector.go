package main

import (
	"math"

	"github.com/shockdev04/altv-go-pkg/alt"
)

func GetForwardVector(rot alt.Vector3) alt.Vector3 {
	z := float64(-rot.Z)
	x := float64(rot.X)
	num := math.Abs(math.Cos(x))
	return alt.Vector3{
		X: float32(-math.Sin(z) * num),
		Y: float32(math.Cos(z) * num),
		Z: float32(math.Sin(x)),
	}
}

func GetPositionInFront(pos alt.Vector3, rot alt.Vector3, dist float32) alt.Vector3 {
	fwd := GetForwardVector(rot)
	return alt.Vector3{
		X: pos.X + fwd.X*dist,
		Y: pos.Y + fwd.Y*dist,
		Z: pos.Z,
	}
}
