package shape

// Circle - структура круга с радиусом и цветом
type Circle struct {
	Shape
	radius float64
	color  string
}

func (c *Circle) Radius() float64 {
	return c.radius
}

func (c *Circle) SetRadius(radius float64) {
	c.radius = radius
}

func (c *Circle) Color() string {
	return c.color
}

func (c *Circle) SetColor(color string) {
	c.color = color
}

func (c *Circle) NameShape() string {
	return "Circle"
}
