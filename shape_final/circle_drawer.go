package shape_final

import "fmt"

// CircleDrawer - структура для рисования круга в консоль
type CircleDrawer struct{}

// DrawShape - метод для рисования круга в консоль
func (cd *CircleDrawer) DrawShape(shape Shape) string {
	// Проверяем, что интерфейс имеет не пустой тип
	if shape == nil {
		return fmt.Sprint("nil shape")
	}

	// Приводим к типу круга
	if s, ok := shape.(*Circle); !ok {
		return fmt.Sprintf("CircleDrawer shape is not a Circle")
	} else {
		return fmt.Sprintf("Circle of radius: %.2f, color: %s", s.Radius(), s.Color())
	}
}
