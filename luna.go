package luna

import (
	"github.com/go-gl/mathgl/mgl64"
)

type Vec2 = mgl64.Vec2
type Vec3 = mgl64.Vec3
type Vec4 = mgl64.Vec4
type Mat4x4 = mgl64.Mat4

func Vector(x, y, z float64) Vec4 {
	return Vec4{x, y, z, 0}
}

func Point(x, y, z float64) Vec4 {
	return Vec4{x, y, z, 1}
}

func Identity() Mat4x4 {
	return Mat4x4{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
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

// TODO: shearing

/*
type material struct {
	color vec3
}

type transform struct {
	m    mat4x4
	inv  mat4x4
	invt mat4x4
}

type interaction struct {
	p    vec4
	n    vec4
	time float64
	obj  shape
}

type ray struct {
	origin    vec4
	direction vec4
}

type shape interface {
	material() *material
	transform() *transform
	normalAt(p vec4) vec4
}

type sphere struct{}

func (s *sphere) material() *material {
	return &material{vec3{0, 0, 0}}
}

func (s *sphere) transform() *transform {
	return &transform{}
}

func (s *sphere) normalAt(p vec4) vec4 {
	return vec4{}
}
*/
