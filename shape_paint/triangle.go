package shape_paint

import "go-training/shape_new"

// TrianglePaint - структура для разукрашивания треугольника
type TrianglePaint struct {
	Triangle *shape_new.Triangle
}

// PaintShape - красим треугольник
func (tp *TrianglePaint) PaintShape(color string) {
	tp.Triangle.SetColor(color)
}
