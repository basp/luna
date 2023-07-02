package luna

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

func (s *Sphere) Intersect(ray *Ray) []*Interaction {
	return []*Interaction{}
}

func (s *Sphere) NormalAt(point Vec4) Vec4 {
	return Vec4{0, 0, 0, 0}
}
