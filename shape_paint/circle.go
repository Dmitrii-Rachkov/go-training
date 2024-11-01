package shape_paint

import (
	"go-training/shape_new"
)

// CirclePaint - структура для разукрашивания круга
type CirclePaint struct {
	Circle *shape_new.Circle
}

// PaintShape - красим круг
func (cp *CirclePaint) PaintShape(color string) {
	cp.Circle.SetColor(color)
}
