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

func TestCrossProduct(t *testing.T) {
	var tests = []struct {
		u    luna.Vec4
		v    luna.Vec4
		want luna.Vec4
	}{
		{
			luna.Vector(1, 2, 3),
			luna.Vector(2, 3, 4),
			luna.Vector(-1, 2, -1),
		},
		{
			luna.Vector(2, 3, 4),
			luna.Vector(1, 2, 3),
			luna.Vector(1, -2, 1),
		},
	}
	for _, tt := range tests {
		ans := luna.Cross(tt.u, tt.v)
		if ans != tt.want {
			t.Errorf("got %v, want %v", ans, tt.want)
		}
	}
}

func TestDotProduct(t *testing.T) {
	var tests = []struct {
		u    luna.Vec4
		v    luna.Vec4
		want float64
	}{
		{
			luna.Vector(1, 2, 3),
			luna.Vector(2, 3, 4),
			20,
		},
		{
			luna.Vector(1, 1, 1),
			luna.Vector(1, 1, 1),
			3,
		},
		{
			luna.Vector(0, 0, 1),
			luna.Vector(1, 1, 1),
			1,
		},
		{
			luna.Vector(0.5, 0.5, 0.5),
			luna.Vector(2, 4, 8),
			7,
		},
	}
	for _, tt := range tests {
		ans := luna.Dot(tt.u, tt.v)
		if ans != tt.want {
			t.Errorf("got %v, want %v", ans, tt.want)
		}
	}
}

func TestHadamardProduct(t *testing.T) {
	var tests = []struct {
		c1   luna.Vec3
		c2   luna.Vec3
		want luna.Vec3
	}{
		{
			luna.Color(1, 0.2, 0.4),
			luna.Color(0.9, 1, 0.1),
			luna.Color(0.9, 0.2, 0.04),
		},
	}
	const threshold = 0.00000001
	for _, tt := range tests {
		ans := luna.Hadamard(tt.c1, tt.c2)
		if !ans.ApproxEqualThreshold(tt.want, threshold) {
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

func TestRayInterpolation(t *testing.T) {
	var tests = []struct {
		time float64
		want luna.Vec4
	}{
		{0, luna.Point(2, 3, 4)},
		{1, luna.Point(3, 3, 4)},
		{-1, luna.Point(1, 3, 4)},
		{2.5, luna.Point(4.5, 3, 4)},
	}
	ray := luna.NewRay(luna.Point(2, 3, 4), luna.Vector(1, 0, 0))
	for _, tt := range tests {
		ans := ray.At(tt.time)
		if !ans.ApproxEqualThreshold(tt.want, 0.000001) {
			t.Errorf("got %v, want %v", ans, tt.want)
		}
	}
}

func TestRaySphereIntersections(t *testing.T) {
	var tests = []struct {
		name string
		ray  luna.Ray
		want []float64
	}{
		{
			"a ray intersects a sphere at two points",
			luna.NewRay(luna.Point(0, 0, -5), luna.Vector(0, 0, 1)),
			[]float64{4.0, 6.0},
		},
		{
			"a ray intersects a sphere at a tangent",
			luna.NewRay(luna.Point(0, 1, -5), luna.Vector(0, 0, 1)),
			[]float64{5.0, 5.0},
		},
		{
			"a ray misses a sphere",
			luna.NewRay(luna.Point(0, 2, -5), luna.Vector(0, 0, 1)),
			[]float64{},
		},
		{
			"a ray originates inside a sphere",
			luna.NewRay(luna.Point(0, 0, 0), luna.Vector(0, 0, 1)),
			[]float64{-1.0, 1.0},
		},
		{
			"a sphere is behind a ray",
			luna.NewRay(luna.Point(0, 0, 5), luna.Vector(0, 0, 1)),
			[]float64{-6.0, -4.0},
		},
	}
	s := luna.NewSphere()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := s.Intersect(tt.ray)
			if len(ans) != len(tt.want) {
				t.Errorf("expected %v intersections", len(tt.want))
			}
			for i := range ans {
				if ans[i].Time != tt.want[i] {
					t.Errorf("got %v, want %v", ans[i].Time, tt.want[i])
				}
			}
		})
	}
}

func TestHit(t *testing.T) {
	p := luna.Point(0, 0, 0)
	n := luna.Vector(0, 0, 0)
	var tests = []struct {
		name string
		xs   []luna.Interaction
		ok   bool
		want float64
	}{
		{
			"all intersections have positive t",
			[]luna.Interaction{
				luna.NewInteraction(p, n, 1, nil),
				luna.NewInteraction(p, n, 2, nil),
			},
			true,
			1,
		},
		{
			"some intersections have negative t",
			[]luna.Interaction{
				luna.NewInteraction(p, n, -1, nil),
				luna.NewInteraction(p, n, 1, nil),
			},
			true,
			1,
		},
		{
			"all intersections have negative t",
			[]luna.Interaction{
				luna.NewInteraction(p, n, -2, nil),
				luna.NewInteraction(p, n, -1, nil),
			},
			false,
			0,
		},
		{
			"hit is always the lowest non-negative intersection",
			[]luna.Interaction{
				luna.NewInteraction(p, n, 5, nil),
				luna.NewInteraction(p, n, 7, nil),
				luna.NewInteraction(p, n, -3, nil),
				luna.NewInteraction(p, n, 2, nil),
			},
			true,
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ok, ans := luna.Hit(tt.xs)
			if ok != tt.ok {
				t.Error("expected to find a hit")
			}
			if ok && ans.Time != tt.want {
				t.Errorf("got %v, want %v", ans.Time, tt.want)
			}
		})
	}
}

func TestRayTransformation(t *testing.T) {
	ray := luna.NewRay(luna.Point(1, 2, 3), luna.Vector(0, 1, 0))
	var tests = []struct {
		name string
		ans  luna.Ray
		want luna.Ray
	}{
		{
			"translating a ray",
			ray.Transform(luna.Translate(3, 4, 5)),
			luna.NewRay(luna.Point(4, 6, 8), luna.Vector(0, 1, 0)),
		},
		{
			"scaling a ray",
			ray.Transform(luna.Scale(2, 3, 4)),
			luna.NewRay(luna.Point(2, 6, 12), luna.Vector(0, 3, 0)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.ans != tt.want {
				t.Errorf("got %v, want %v", tt.ans, tt.want)
			}
		})
	}
}

func TestIntersectTransformedSphere(t *testing.T) {
	var tests = []struct {
		name      string
		transform *luna.Transform
		want      []float64
	}{
		{
			"intersect a scaled sphere with a ray",
			luna.NewTransform(luna.Scale(2, 2, 2)),
			[]float64{3, 7},
		},
		{
			"intersect a translated sphere with a ray",
			luna.NewTransform(luna.Translate(5, 0, 0)),
			[]float64{},
		},
	}
	s := luna.NewSphere()
	r := luna.NewRay(luna.Point(0, 0, -5), luna.Vector(0, 0, 1))
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s.SetTransform(tt.transform)
			ans := s.Intersect(r)
			if len(ans) != len(tt.want) {
				t.Fatalf("Expected %v intersections", len(tt.want))
			}
			for i := range ans {
				if ans[i].Time != tt.want[i] {
					t.Errorf("got %v, want %v", ans[i].Time, tt.want[i])
				}
			}
		})
	}
}

func TestSphereNormalCalculation(t *testing.T) {
	var tests = []struct {
		name string
		p    luna.Vec4
		want luna.Vec4
	}{
		{
			"the normal on a sphere at a point on the x axis",
			luna.Point(1, 0, 0),
			luna.Vector(1, 0, 0),
		},
		{
			"the normal on a sphere at a point on the y axis",
			luna.Point(0, 1, 0),
			luna.Vector(0, 1, 0),
		},
		{
			"the normal on a sphere at a point on the z axis",
			luna.Point(0, 0, 1),
			luna.Vector(0, 0, 1),
		},
		{
			"the normal on a sphere at a non-axial point",
			luna.Point(Sqrt3Over3, Sqrt3Over3, Sqrt3Over3),
			luna.Vector(Sqrt3Over3, Sqrt3Over3, Sqrt3Over3),
		},
	}
	s := luna.NewSphere()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := s.NormalAt(tt.p)
			if !ans.ApproxEqualThreshold(tt.want, 0.0000001) {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}

func TestNormalAtIsNormalized(t *testing.T) {
	s := luna.NewSphere()
	n := s.NormalAt(luna.Point(Sqrt3Over3, Sqrt3Over3, Sqrt3Over3))
	if n != n.Normalize() {
		t.Error("normal should be normalized")
	}
}
