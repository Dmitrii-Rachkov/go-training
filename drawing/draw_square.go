package drawing

import "fmt"

// Square - структура квадрата со стороной и цветом
type Square struct {
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

// Draw - рисуем квадрат
func (s *Square) Draw() string {
	return fmt.Sprintf("Square drawn with side: %.2f and color: %s", s.side, s.color)

}
