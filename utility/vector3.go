package utility

import (
	"math"
	"math/rand"

	"github.com/timo972/altv-go-pkg/alt"
)

func RandomPositionAround(pos alt.Vector3, minRange float32, maxRange float32) alt.Vector3 {
	pos.X = pos.X + (rand.Float32()*(maxRange-minRange+1) + minRange)
	pos.Y = pos.Y + (rand.Float32()*(maxRange-minRange+1) + minRange)

	return pos
}

func ForwardVector(rot alt.Vector3) alt.Vector3 {
	z := float64(-rot.Z)
	x := float64(rot.X)
	num := math.Abs(math.Cos(x))
	return alt.Vector3{
		X: float32(-math.Sin(z) * num),
		Y: float32(math.Cos(z) * num),
		Z: float32(math.Sin(x)),
	}
}

func PositionInFront(pos alt.Vector3, rot alt.Vector3, dist float32) alt.Vector3 {
	fwd := ForwardVector(rot)
	return alt.Vector3{
		X: pos.X + fwd.X*dist,
		Y: pos.Y + fwd.Y*dist,
		Z: pos.Z,
	}
}
