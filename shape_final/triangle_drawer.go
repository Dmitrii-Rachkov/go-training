package shape_final

import "fmt"

// TriangleDrawer - структура для рисования треугольника
type TriangleDrawer struct{}

// DrawShape - метод для рисования треугольника в консоль
func (td *TriangleDrawer) DrawShape(shape Shape) string {
	// Проверяем, что интерфейс имеет не пустой тип
	if shape == nil {
		return fmt.Sprint("nil shape")
	}

	// Приводим к типу квадрата
	if s, ok := shape.(*Triangle); !ok {
		return fmt.Sprintf("TriangleDrawer shape is not a Triangle")
	} else {
		return fmt.Sprintf("Triangle base: %.2f, height: %.2f, color: %s", s.Base(), s.Height(), s.Color())
	}
}
