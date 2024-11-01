package drawing

import "fmt"

// Circle - структура круга с радиусом и цветом
type Circle struct {
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

// Draw - рисуем круг
func (c *Circle) Draw() string {
	return fmt.Sprintf("Circle drawn with radius: %.2f and color: %s", c.radius, c.color)
}
