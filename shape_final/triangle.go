package shape_final

import "fmt"

// Triangle - структура треугольника с основаниями и цветом
type Triangle struct {
	base   float64
	height float64
	color  string
}

// NewTriangle - конструктор для создания нового треугольника
func NewTriangle(base, height float64, color string) *Triangle {
	return &Triangle{base: base, height: height, color: color}
}

// Area - метод для вычисления площади треугольника
func (t *Triangle) Area() float64 {
	return 0.5 * t.base * t.height
}

// Color - метод для получения цвета треугольника
func (t *Triangle) Color() string {
	return t.color
}

// SetColor - метод для установки цвета треугольника
func (t *Triangle) SetColor(color string) {
	t.color = color
}

// Base - метод для получения основания треугольника
func (t *Triangle) Base() float64 {
	return t.base
}

// SetBase - метод для установки основания треугольника
func (t *Triangle) SetBase(base float64) {
	t.base = base
}

// Height - метод для получения высоты треугольника
func (t *Triangle) Height() float64 {
	return t.height
}

// SetHeight - метод для установки высоты треугольника
func (t *Triangle) SetHeight(height float64) {
	t.height = height
}

// DrawShape - метод для рисования треугольника в консоль
func (t *Triangle) DrawShape() string {
	return fmt.Sprintf("Triangle with base: %.2f, height: %.2f, color: %s", t.Base(), t.Height(), t.Color())
}

// Paint - метод для раскрашивания треугольника
func (t *Triangle) Paint(color string) {
	t.SetColor(color)
}
