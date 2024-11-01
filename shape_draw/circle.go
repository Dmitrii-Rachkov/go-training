package shape_draw

import (
	"fmt"
	"go-training/shape_new"
)

// CircleDraw - структура для рисования круга
type CircleDraw struct {
	Circle *shape_new.Circle
}

// DrawShape - рисуем круг
func (cd *CircleDraw) DrawShape() string {
	return fmt.Sprintf("Circle of radius: %.2f", cd.Circle.Radius())
}
