package luna

type Transform struct {
	M    Mat4x4
	Inv  Mat4x4
	InvT Mat4x4
}

func NewTransform(m Mat4x4) *Transform {
	return &Transform{m, m.Inv(), m.Inv().Transpose()}
}
