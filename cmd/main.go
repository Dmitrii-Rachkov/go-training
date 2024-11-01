package main

import (
	"fmt"
	"go-training/shape_final"
)

func main() {
	// ФИНАЛЬНАЯ ВЕРСИЯ

	// Создание фигур
	circle := shape_final.NewCircle(5.0, "red")
	square := shape_final.NewSquare(4.0, "blue")
	triangle := shape_final.NewTriangle(3.0, 4.0, "green")

	//// Создание рисовальщика
	//drawer := &shape_new.ConsoleDrawer{}

	// Рисование фигур до раскрашивания
	//fmt.Println(drawer.DrawShape(circle))
	//fmt.Println(drawer.DrawShape(square))
	//fmt.Println(drawer.DrawShape(triangle))
	fmt.Println(circle.DrawShape())
	fmt.Println(square.DrawShape())
	fmt.Println(triangle.DrawShape())

	// Раскрашивание фигур
	circle.Paint("yellow")
	square.Paint("purple")
	triangle.Paint("orange")

	fmt.Println(circle.Color())
	fmt.Println(square.Color())
	fmt.Println(triangle.Color())
}
