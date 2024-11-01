package shape

// Square - структура квадрата со стороной и цветом
type Square struct {
	Shape
	side  float64
	color string
}

func (s *Square) Side() float64 {
	return s.side
}

func (s *Square) SetSide(side float64) {
	s.side = side
}

func (s *Square) Color() string {
	return s.color
}

func (s *Square) SetColor(color string) {
	s.color = color
}

func (s *Square) NameShape() string {
	return "Square"
}
