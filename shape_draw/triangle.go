package shape_draw

import (
	"fmt"
	"go-training/shape_new"
)

// TriangleDraw - структура для рисования треугольника
type TriangleDraw struct {
	Triangle shape_new.Triangle
}

// DrawShape - рисуем треугольник
func (td *TriangleDraw) DrawShape() string {
	return fmt.Sprintf("Triangle of base: %.2f sideA: %.2f sideB: %.2f",
		td.Triangle.Base(), td.Triangle.SideA(), td.Triangle.SideB())
}
