package luna_test

import (
	"math"
	"testing"

	"github.com/basp/luna"
)

const PiOver4 = math.Pi / 4
const PiOver2 = math.Pi / 2
const Sqrt2Over2 = math.Sqrt2 / 2

var Sqrt3Over3 = math.Sqrt(3) / 3

func TestIdentity(t *testing.T) {
	var tests = []struct {
		row  int
		col  int
		want float64
	}{
		{0, 0, 1},
		{0, 1, 0},
		{0, 2, 0},
		{0, 3, 0},
		{1, 0, 0},
		{1, 1, 1},
		{1, 2, 0},
		{1, 3, 0},
		{2, 0, 0},
		{2, 1, 0},
		{2, 2, 1},
		{2, 3, 0},
		{3, 0, 0},
		{3, 1, 0},
		{3, 2, 0},
		{3, 3, 1},
	}
	m := luna.Identity()
	for _, tt := range tests {
		i := m.Index(tt.row, tt.col)
		ans := m[i]
		if ans != tt.want {
			t.Errorf("got %v, want %v", ans, tt.want)
		}
	}
}

func TestPrimitiveTransformations(t *testing.T) {
	var tests = []struct {
		name      string
		transform luna.Mat4x4
		v         luna.Vec4
		want      luna.Vec4
	}{
		{
			"multiplying by a translation matrix",
			luna.Translate(5, -3, 2),
			luna.Point(-3, 4, 5),
			luna.Point(2, 1, 7),
		},
		{
			"multiplying by the inverse of a translation matrix",
			luna.Translate(5, -3, 2).Inv(),
			luna.Point(-3, 4, 5),
			luna.Point(-8, 7, 3),
		},
		{
			"translation does not affect vectors",
			luna.Translate(5, 3, 2),
			luna.Vector(-3, 4, 5),
			luna.Vector(-3, 4, 5),
		},
		{
			"scaling matrix applied to a point",
			luna.Scale(2, 3, 4),
			luna.Point(-4, 6, 8),
			luna.Point(-8, 18, 32),
		},
		{
			"scaling matrix applied to a vector",
			luna.Scale(2, 3, 4),
			luna.Vector(-4, 6, 8),
			luna.Vector(-8, 18, 32),
		},
		{
			"multiplying by the inverse of a scaling matrix",
			luna.Scale(2, 3, 4).Inv(),
			luna.Vector(-4, 6, 8),
			luna.Vector(-2, 2, 2),
		},
		{
			"reflection is scaling by a negative value",
			luna.Scale(-1, 1, 1),
			luna.Point(2, 3, 4),
			luna.Point(-2, 3, 4),
		},
		{
			"rotating a point 45 degrees around the x axis",
			luna.RotateX(PiOver4),
			luna.Point(0, 1, 0),
			luna.Point(0, Sqrt2Over2, Sqrt2Over2),
		},
		{
			"rotating a point 90 degrees around the x axis",
			luna.RotateX(PiOver2),
			luna.Point(0, 1, 0),
			luna.Point(0, 0, 1),
		},
		{
			"the inverse of an x-rotation rotates in the opposite direction",
			luna.RotateX(PiOver4).Inv(),
			luna.Point(0, 1, 0),
			luna.Point(0, Sqrt2Over2, -Sqrt2Over2),
		},
		{
			"rotating a point 45 degrees around the y axis",
			luna.RotateY(PiOver4),
			luna.Point(0, 0, 1),
			luna.Point(Sqrt2Over2, 0, Sqrt2Over2),
		},
		{
			"rotating a point 90 degrees around the y axis",
			luna.RotateY(PiOver2),
			luna.Point(0, 0, 1),
			luna.Point(1, 0, 0),
		},
		{
			"rotating a point 45 degrees around the z axis",
			luna.RotateZ(PiOver4),
			luna.Point(0, 1, 0),
			luna.Point(-Sqrt2Over2, Sqrt2Over2, 0),
		},
		{
			"rotating a point 90 degrees around the z axis",
			luna.RotateZ(PiOver2),
			luna.Point(0, 1, 0),
			luna.Point(-1, 0, 0),
		},
	}
	const threshold = 0.000001
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := tt.transform.Mul4x1(tt.v)
			if !ans.ApproxEqualThreshold(tt.want, threshold) {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
