package drawing

// DrawMethod - интерфейс метода рисования
type DrawMethod interface {
	Draw() string
}

// Drawer - структура открытая для расширения и закрытая для изменения
// При использовании DrawMethod не нужно редактировать Drawer поведение при добавлении новых способов рисования
type Drawer struct{}

// Process - действие по рисованию
func (d Drawer) Process(dm DrawMethod) string {
	return dm.Draw()
}
