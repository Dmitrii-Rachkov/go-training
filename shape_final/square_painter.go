package shape_final

import "fmt"

// SquarePainter - структура для раскрашивания квадрата
type SquarePainter struct{}

// Paint - метод для раскрашивания круга в консоль
func (sp *SquarePainter) Paint(color string, shape Shape) string {
	// Проверяем, что интерфейс имеет не пустой тип
	if shape == nil {
		return fmt.Sprint("nil shape")
	}

	// Приводим к типу квадрата
	if s, ok := shape.(*Square); !ok {
		return fmt.Sprintf("SquarePainter shape is not a Square")
	} else {
		s.SetColor(color)
		return fmt.Sprintf("Square of color: %s", s.Color())
	}
}
