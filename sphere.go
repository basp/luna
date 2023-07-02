package luna

import "math"

type Sphere struct {
	material  *Material
	transform *Transform
}

func NewSphere() *Sphere {
	return &Sphere{
		NewDefaultMaterial(),
		NewImplicitTransform(Identity()),
	}
}

func (s *Sphere) Material() *Material {
	return s.material
}

func (s *Sphere) Intersect(ray Ray) []Interaction {
	ray = ray.Transform(s.transform.Inv)
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

func (s *Sphere) NormalAt(wp Vec4) Vec4 {
	op := s.transform.Inv.Mul4x1(wp)
	on := op.Sub(Point(0, 0, 0))
	wn := s.transform.InvT.Mul4x1(on)
	wn[3] = 0
	return wn.Normalize()
}

func (s *Sphere) SetTransform(t *Transform) {
	s.transform = t
}
