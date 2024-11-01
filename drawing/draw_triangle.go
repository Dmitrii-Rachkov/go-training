package drawing

import (
	"fmt"
	"go-training/shape"
)

// TriangleDraw - встраиваем структуру треугольника
type TriangleDraw struct {
	shape.Triangle
}

// Draw - рисуем треугольник
func (t *TriangleDraw) Draw() string {
	return fmt.Sprintf("Triangle drawn with base: %.2f sideA: %.2f sideB: %.2f and color: %s",
		t.Base(), t.SideA(), t.SideB(), t.Color())
}
