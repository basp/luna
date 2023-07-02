package luna

import (
	"github.com/go-gl/mathgl/mgl64"
)

type Vec2 = mgl64.Vec2
type Vec3 = mgl64.Vec3
type Vec4 = mgl64.Vec4
type Mat4x4 = mgl64.Mat4

type Material struct {
	Color   Vec3
	Ambient float64
}

type Transform struct {
	M    Mat4x4
	Inv  Mat4x4
	InvT Mat4x4
}

type Shape interface {
	Material() *Material
	Intersect(ray *Ray) []*Interaction
	NormalAt(point Vec4) Vec4
}

type Interaction struct {
	Point  Vec4
	Normal Vec4
	Time   float64
	Object Shape
}

type isects []*Interaction

func (xs isects) Len() int           { return len(xs) }
func (xs isects) Swap(i, j int)      { xs[i], xs[j] = xs[j], xs[i] }
func (xs isects) Less(i, j int) bool { return xs[i].Time < xs[j].Time }

func Vector(x, y, z float64) Vec4 {
	return Vec4{x, y, z, 0}
}

func Point(x, y, z float64) Vec4 {
	return Vec4{x, y, z, 1}
}

func Color(r, g, b float64) Vec3 {
	return Vec3{r, g, b}
}

func Identity() Mat4x4 {
	return Mat4x4{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
}

func Dot(u, v Vec4) float64 {
	return u.Dot(v)
}

func Cross(u, v Vec4) Vec4 {
	return u.Vec3().Cross(v.Vec3()).Vec4(0)
}

func Hadamard(u, v Vec3) Vec3 {
	return Vec3{
		u[0] * v[0],
		u[1] * v[1],
		u[2] * v[2],
	}
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

func NewInteraction(p, n Vec4, t float64, obj Shape) Interaction {
	return Interaction{p, n, t, obj}
}

func NewTransform(m Mat4x4) *Transform {
	return &Transform{m, m.Inv(), m.Inv().Transpose()}
}

func NewDefaultMaterial() *Material {
	return &Material{}
}
