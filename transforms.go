package luna

type Transform struct {
	M    Mat4x4
	Inv  Mat4x4
	InvT Mat4x4
}

func NewImplicitTransform(m Mat4x4) *Transform {
	return &Transform{m, m.Inv(), m.Inv().Transpose()}
}

func NewExplicitTransform(m Mat4x4, inv Mat4x4) *Transform {
	return &Transform{m, inv, inv.Transpose()}
}
