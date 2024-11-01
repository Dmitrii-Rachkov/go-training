package shape_paint

import "go-training/shape_new"

// SquarePaint - структура для разукрашивания квадрата
type SquarePaint struct {
	Square *shape_new.Square
}

// PaintShape - красим квадрат
func (sp *SquarePaint) PaintShape(color string) {
	sp.Square.SetColor(color)
}
