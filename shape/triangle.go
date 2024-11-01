package shape

// Triangle - структура треугольника с основанием, двумя сторонами и цветом
type Triangle struct {
	Shape
	base  float64
	sideA float64
	sideB float64
	color string
}

func (t *Triangle) Base() float64 {
	return t.base
}

func (t *Triangle) SetBase(base float64) {
	t.base = base
}

func (t *Triangle) SideA() float64 {
	return t.sideA
}

func (t *Triangle) SetSideA(sideA float64) {
	t.sideA = sideA
}

func (t *Triangle) SideB() float64 {
	return t.sideB
}

func (t *Triangle) SetSideB(sideB float64) {
	t.sideB = sideB
}

func (t *Triangle) Color() string {
	return t.color
}

func (t *Triangle) SetColor(color string) {
	t.color = color
}

func (t *Triangle) NameShape() string {
	return "Triangle"
}
