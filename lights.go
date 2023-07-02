package luna

type PointLight struct {
	Position  Vec4
	Intensity Vec3
}

func NewPointLight(p Vec4, i Vec3) *PointLight {
	return &PointLight{p, i}
}
