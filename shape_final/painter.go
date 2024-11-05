package shape_final

// Painter - интерфейс для раскрашивания фигур
type Painter interface {
	Paint(color string, shape Shape) string
}
