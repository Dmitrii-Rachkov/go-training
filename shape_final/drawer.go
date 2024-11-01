package shape_final

// ShapeDrawer - интерфейс для рисования фигур
type ShapeDrawer interface {
	DrawShape() string
}

//// ConsoleDrawer - структура для рисования фигур в консоль
//type ConsoleDrawer struct{}

//// DrawShape - метод для рисования фигур в консоль
//func (cd *ConsoleDrawer) DrawShape(shape Shape) string {
//	switch s := shape.(type) {
//	case *Circle:
//		return fmt.Sprintf("Circle of radius: %.2f, color: %s", s.Radius(), s.Color())
//	case *Square:
//		return fmt.Sprintf("Square of side: %.2f, color: %s", s.Side(), s.Color())
//	default:
//		return "Unknown shape"
//	}
//}
