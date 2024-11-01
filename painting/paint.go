package painting

// PaintMethod - интерфейс метода раскрашивания
type PaintMethod interface {
	Paint() string
}

// Painter - структура открытая для расширения и закрытая для изменения
// При использовании PaintMethod не нужно редактировать Painter поведение при добавлении новых способов раскрашивания
type Painter struct{}

// Process - действие по рисованию
func (p Painter) Process(pm PaintMethod) string {
	return pm.Paint()
}
