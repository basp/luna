package luna

import "math"

type Sphere struct {
	material  *Material
	transform *Transform
}

func NewSphere() *Sphere {
	return new(Sphere)
}

func (s *Sphere) Material() *Material {
	return s.material
}

func (s *Sphere) Intersect(ray Ray) []Interaction {
	sphereToRay := ray.Origin.Sub(Point(0, 0, 0))
	a := Dot(ray.Direction, ray.Direction)
	b := 2 * Dot(ray.Direction, sphereToRay)
	c := Dot(sphereToRay, sphereToRay) - 1
	d := b*b - 4*a*c
	if d < 0 {
		return []Interaction{}
	}
	t0 := (-b - math.Sqrt(d)) / (2 * a)
	t1 := (-b + math.Sqrt(d)) / (2 * a)
	p0 := ray.At(t0)
	p1 := ray.At(t1)
	n0 := s.NormalAt(p0)
	n1 := s.NormalAt(p1)
	return []Interaction{
		NewInteraction(p0, n0, t0, s),
		NewInteraction(p1, n1, t1, s),
	}
}

func (s *Sphere) NormalAt(p Vec4) Vec4 {
	return Vec4{0, 0, 0, 0}
}
