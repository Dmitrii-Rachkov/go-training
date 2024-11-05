package shape_final

import "fmt"

// CirclePainter - структура для раскрашивания круга
type CirclePainter struct{}

// Paint - метод для раскрашивания круга в консоль
func (cp *CirclePainter) Paint(color string, shape Shape) string {
	// Проверяем, что интерфейс имеет не пустой тип
	if shape == nil {
		return fmt.Sprint("nil shape")
	}

	// Приводим к типу круга
	if s, ok := shape.(*Circle); !ok {
		return fmt.Sprintf("CirclePainter shape is not a Circle")
	} else {
		s.SetColor(color)
		return fmt.Sprintf("Circle of color: %s", s.Color())
	}
}
