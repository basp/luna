package luna

import (
	"sort"

	"github.com/go-gl/mathgl/mgl64"
)

type Vec2 = mgl64.Vec2
type Vec3 = mgl64.Vec3
type Vec4 = mgl64.Vec4
type Mat4x4 = mgl64.Mat4

type Shape interface {
	Material() *Material
	Intersect(ray Ray) []Interaction
	NormalAt(point Vec4) Vec4
	SetTransform(t *Transform)
}

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

func Hit(xs isects) (bool, *Interaction) {
	sort.Sort(xs)
	for _, x := range xs {
		if x.Time >= 0 {
			return true, &x
		}
	}
	return false, nil
}

func Reflect(v, n Vec4) Vec4 {
	return v.Sub(n.Mul(2).Mul(Dot(v, n)))
}
