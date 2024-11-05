package main

import (
	"fmt"
	shapeFinal "go-training/shape_final"
)

func main() {
	// ФИНАЛЬНАЯ ВЕРСИЯ

	// Создание фигур
	circle := shapeFinal.NewCircle(5.0, "red")
	square := shapeFinal.NewSquare(4.0, "blue")
	triangle := shapeFinal.NewTriangle(3.0, 4.0, "green")

	// Создаём структуры рисовальщиков для наших фигур
	cd := shapeFinal.CircleDrawer{}
	sd := shapeFinal.SquareDrawer{}
	td := shapeFinal.TriangleDrawer{}

	// Рисуем фигуры
	fmt.Println(cd.DrawShape(circle))
	fmt.Println(sd.DrawShape(square))
	fmt.Println(td.DrawShape(triangle))

	// Создаём структуры для раскрашивания фигур
	cp := shapeFinal.CirclePainter{}
	sp := shapeFinal.SquarePainter{}
	tp := shapeFinal.TrianglePainter{}

	// Раскрашиваем фигуры
	fmt.Println(cp.Paint("purple", circle))
	fmt.Println(sp.Paint("orange", square))
	fmt.Println(tp.Paint("yellow", triangle))
}
