package shape_final

// Circle - структура круга с радиусом и цветом
type Circle struct {
	radius float64
	color  string
}

// NewCircle - конструктор для создания нового круга
func NewCircle(radius float64, color string) *Circle {
	return &Circle{radius: radius, color: color}
}

// Radius - метод для получения радиуса круга
func (c *Circle) Radius() float64 {
	return c.radius
}

// SetRadius - метод для установки радиуса круга
func (c *Circle) SetRadius(radius float64) {
	c.radius = radius
}

// Area - метод для вычисления площади круга
func (c *Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

// Color - метод для получения цвета круга
func (c *Circle) Color() string {
	return c.color
}

// SetColor - метод для установки цвета круга
func (c *Circle) SetColor(color string) {
	c.color = color
}
