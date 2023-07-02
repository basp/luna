package luna

type Material struct {
	Color   Vec3
	Ambient float64
}

func NewDefaultMaterial() *Material {
	return &Material{
		Vec3{1, 1, 1},
		0.1,
	}
}
