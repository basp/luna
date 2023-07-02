package luna

type Ray struct {
	Origin    Vec4
	Direction Vec4
}

func NewRay(origin, direction Vec4) Ray {
	return Ray{origin, direction}
}

func (ray Ray) At(time float64) Vec4 {
	return ray.Origin.Add(ray.Direction.Mul(time))
}

func (ray Ray) Transform(m Mat4x4) Ray {
	return Ray{m.Mul4x1(ray.Origin), m.Mul4x1(ray.Direction)}
}
