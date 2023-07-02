package luna_test

import (
	"testing"

	"github.com/basp/luna"
)

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
	m := luna.Identity
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
	}
	for _, tt := range tests {
		ans := tt.transform.Mul4x1(tt.v)
		if ans != tt.want {
			t.Errorf("got %v, want %v", ans, tt.want)
		}
	}
}
