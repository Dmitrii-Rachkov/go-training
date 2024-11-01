package shape_draw

import (
	"fmt"
	"go-training/shape_new"
)

// SquareDraw - структура для рисования квадрата
type SquareDraw struct {
	Square shape_new.Square
}

// DrawShape - рисуем квадрат
func (sd *SquareDraw) DrawShape() string {
	return fmt.Sprintf("Square of side: %.2f", sd.Square.Side())
}
