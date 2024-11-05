package shape_final

// Shape - интерфейс для всех фигур
type Shape interface {
	Area() float64
	Color() string
	SetColor(color string)
	// GetData - возвращает стороны и другие параметры
}
