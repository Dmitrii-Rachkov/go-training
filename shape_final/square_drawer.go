package shape_final

import "fmt"

// SquareDrawer - структура для рисования квадрата
type SquareDrawer struct{}

// DrawShape - метод для рисования круга в консоль
func (sd *SquareDrawer) DrawShape(shape Shape) string {
	// Проверяем, что интерфейс имеет не пустой тип
	if shape == nil {
		return fmt.Sprint("nil shape")
	}

	// Приводим к типу квадрата
	if s, ok := shape.(*Square); !ok {
		return fmt.Sprintf("SquareDrawer shape is not a Square")
	} else {
		return fmt.Sprintf("Square of side: %.2f, color: %s", s.Side(), s.Color())
	}
}
