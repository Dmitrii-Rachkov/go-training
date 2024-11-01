package main

import (
	"fmt"
	"go-training/shape_final"
)

func main() {
	//// Создаём рисовальщик
	//drawer := drawing.Drawer{}
	//
	//// Создаём круг и его параметры
	//circle := drawing.Circle{}
	//circle.SetRadius(8.5)
	//circle.SetColor("red")
	//
	//// Рисуем круг
	//fmt.Println(drawer.Process(&circle))
	//
	//// Создаём квадрат и его параметры
	//square := drawing.Square{}
	//square.SetSide(5.2)
	//square.SetColor("blue")
	//
	//// Рисуем квадрат
	//fmt.Println(drawer.Process(&square))
	//
	//// Создаём треугольник и его параметры
	//triangle := shape.Triangle{}
	//triangle.SetBase(5.5)
	//triangle.SetSideA(6.0)
	//triangle.SetSideB(7.8)
	//triangle.SetColor("yellow")
	//
	//// Рисуем треугольник
	//fmt.Println(drawer.Process(triangle))

	// НОВАЯ ВЕРСИЯ

	// Структура рисования
	// d := shape_draw.Drawing{}

	//// Создаём круг
	//circle := shape_new.Circle{}
	//circle.SetRadius(8.3)
	//
	//// Рисуем круг
	//cd := shape_draw.CircleDraw{Circle: &circle}
	//// Не удаётся получить строку из функции d.Process(&cd), она работает и заходит
	//// в функцию DrawShape(), но ничего не возвращает
	//// d.Process(&cd)
	//fmt.Println(cd.DrawShape())
	//
	//// Красим круг
	//cp := shape_paint.CirclePaint{Circle: &circle}
	//cp.PaintShape("red")
	//fmt.Println(circle.Color())
	//
	//// Создаём квадрат
	//square := shape_new.Square{}
	//square.SetSide(5.0)
	//
	//// Рисуем квадрат
	//sd := shape_draw.SquareDraw{Square: square}
	//fmt.Println(sd.DrawShape())
	//
	//// Красим квадрат
	//sp := shape_paint.SquarePaint{Square: &square}
	//sp.PaintShape("blue")
	//fmt.Println(square.Color())
	//
	//// Создаём треугольник
	//triangle := shape_new.Triangle{}
	//triangle.SetBase(4.9)
	//triangle.SetSideA(7.5)
	//triangle.SetSideB(9.8)
	//
	//// Рисуем треугольник
	//td := shape_draw.TriangleDraw{Triangle: triangle}
	//fmt.Println(td.DrawShape())
	//
	//// Красим треугольник
	//tp := shape_paint.TrianglePaint{Triangle: &triangle}
	//tp.PaintShape("yellow")
	//fmt.Println(triangle.Color())

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
