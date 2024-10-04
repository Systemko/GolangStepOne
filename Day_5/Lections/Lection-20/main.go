package main

import (
	"fmt"
	"math"
)

/*
	Структура - явно декларированный заименованный набор состояний.
	Структура, отвечает на вопрос - из чего я должна состоять,
	что-бы считаться тем ТИПОМ, который декларируется структурой.
	Структуры описывает ПАТТЕРН СОСТОЯНИЯ.

	Интерфейс - явно декларированный набор сигнатур ПОВЕДЕНИЙ
	(чаще всего в виде набора методов), удовлетворив который,
	может считаться ТИПОМ, который декларирует интерфейс.
	Интерфейс, отвечает на вопрос - что я должен уметь,
	что-бы считаться тем ТИПОМ, который декларируется интерфейсом.
	Интерфейс описывет ПАТТЕРН ПОВЕДЕНИЯ.
*/

// Объявление интерфейса
type FigureBuilder interface {
	// Набор сигнатур методов, которые
	// необходимо реализовать в структуре претенденте.
	Area() int
	Perimeter() int
}

// Декларируем структуры претендентов
// Когда методы Area() и Perimeter() реализованы,
// говорят, что Rectangle удовлетворяет условиям интерфейса FigureBuilder
type Rectangle struct {
	width, length int
}

func (r Rectangle) Area() int {
	return r.width * r.length
}

func (r Rectangle) Perimeter() int {
	return 2 * (r.width + r.length)
}

type Circle struct {
	radius int
}

func (c Circle) Area() int {
	return int(math.Round(2 * math.Pi * float64(c.radius)))
}

func (c Circle) Perimeter() int {
	return int(math.Round(math.Pi * float64(c.radius*c.radius)))
}

func main() {
	// Создадим по 3 экземпляра круга и прямоугольника
	// и рассчитаем общую площадь.
	r1 := Rectangle{10, 20}
	r2 := Rectangle{30, 50}
	r3 := Rectangle{1, 1}
	c1 := Circle{5}
	c2 := Circle{10}
	c3 := Circle{15}

	rectangles := []Rectangle{r1, r2, r3}
	totalAreaOfRectangles := 0
	for _, rect := range rectangles {
		totalAreaOfRectangles += rect.Area()
	}

	circles := []Circle{c1, c2, c3}
	totalAreaOfCircles := 0
	for _, circle := range circles {
		totalAreaOfCircles += circle.Area()
	}

	fmt.Println("Total area is:", totalAreaOfRectangles+totalAreaOfCircles)

	// Теперь воспользуемся интерфейсом, который мы создали выше:
	figures := []FigureBuilder{r1, r2, r3, c1, c2, c3}
	totalArea := 0
	for _, figure := range figures {
		totalArea += figure.Area()
	}
	fmt.Println("Total area is:", totalArea)

}
