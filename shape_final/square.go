package shape_final

// Square - структура квадрата со стороной и цветом
type Square struct {
	side  float64
	color string
}

// NewSquare - конструктор для создания нового квадрата
func NewSquare(side float64, color string) *Square {
	return &Square{side: side, color: color}
}

// Area - метод для вычисления площади квадрата
func (s *Square) Area() float64 {
	return s.side * s.side
}

// Color - метод для получения цвета квадрата
func (s *Square) Color() string {
	return s.color
}

// SetColor - метод для установки цвета квадрата
func (s *Square) SetColor(color string) {
	s.color = color
}

// Side - метод для получения стороны квадрата
func (s *Square) Side() float64 {
	return s.side
}

// SetSide - метод для установки стороны квадрата
func (s *Square) SetSide(side float64) {
	s.side = side
}
