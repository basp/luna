package luna

import (
	"github.com/go-gl/mathgl/mgl64"
)

type Vec3 = mgl64.Vec3
type Vec4 = mgl64.Vec4
type Mat4x4 = mgl64.Mat4

var Identity = &Mat4x4{
	1, 0, 0, 0,
	0, 1, 0, 0,
	0, 0, 1, 0,
	0, 0, 0, 1,
}

func Vector(x, y, z float64) Vec4 {
	return Vec4{x, y, z, 0}
}

func Point(x, y, z float64) Vec4 {
	return Vec4{x, y, z, 1}
}

func Translate(tx, ty, tz float64) Mat4x4 {
	return mgl64.Translate3D(tx, ty, tz)
}

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
