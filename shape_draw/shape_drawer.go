package shape_draw

// ShapeDrawer - интерфейс для рисования фигур
type ShapeDrawer interface {
	DrawShape() string
}

// Drawing и Process - попытка реализации принципа открытости/закрытости
type Drawing struct{}

func (p Drawing) Process(sd ShapeDrawer) {
	sd.DrawShape()
}
