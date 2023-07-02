package luna

import "github.com/go-gl/mathgl/mgl64"

type Interaction struct {
	Point  Vec4
	Normal Vec4
	Time   float64
	Object Shape
}

type isects []Interaction

func (xs isects) Len() int           { return len(xs) }
func (xs isects) Swap(i, j int)      { xs[i], xs[j] = xs[j], xs[i] }
func (xs isects) Less(i, j int) bool { return xs[i].Time < xs[j].Time }

func NewInteraction(p, n Vec4, t float64, obj Shape) Interaction {
	return Interaction{p, n, t, obj}
}

func Translate(tx, ty, tz float64) Mat4x4 {
	return mgl64.Translate3D(tx, ty, tz)
}

func Scale(sx, sy, sz float64) Mat4x4 {
	return mgl64.Scale3D(sx, sy, sz)
}

func RotateX(angle float64) Mat4x4 {
	return mgl64.HomogRotate3DX(angle)
}

func RotateY(angle float64) Mat4x4 {
	return mgl64.HomogRotate3DY(angle)
}

func RotateZ(angle float64) Mat4x4 {
	return mgl64.HomogRotate3DZ(angle)
}
