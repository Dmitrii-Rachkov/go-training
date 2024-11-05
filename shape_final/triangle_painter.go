package shape_final

import "fmt"

// TrianglePainter - структура для раскрашивания треугольника
type TrianglePainter struct{}

// Paint - метод для раскрашивания круга в консоль
func (tp *TrianglePainter) Paint(color string, shape Shape) string {
	// Проверяем, что интерфейс имеет не пустой тип
	if shape == nil {
		return fmt.Sprint("nil shape")
	}

	// Приводим к типу треугольника
	if s, ok := shape.(*Triangle); !ok {
		return fmt.Sprintf("TrianglePainter shape is not a Triangle")
	} else {
		s.SetColor(color)
		return fmt.Sprintf("Triangle of color: %s", s.Color())
	}
}
