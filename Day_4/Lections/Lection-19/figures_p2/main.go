package main

import (
	"fmt"
	"math"
)

// Преимущество методов над функциями
// - Улучшает читаемость и консистентность кода
// - Запрещается создание функций с одинаковыми названиями, а методы
// для различных структур могут иметь одинаковые имена

type Circle struct {
	radius float64
}

type Rectangle struct {
	lenght, width int
}

// Создадим функцию и метод Perimeter для структуры Circle
func PerimeterCircle(c Circle) float64 {
	return 2 * c.radius * math.Pi
}

func (c Circle) Perimeter() float64 {
	return 2 * c.radius * math.Pi
}

// Создадим функцию и метод Perimeter для структуры Rectangle
func PerimeterRectangle(r Rectangle) int {
	return 2 * (r.lenght + r.width)
}

func (r Rectangle) Perimeter() int {
	return 2 * (r.lenght + r.width)
}

func main() {
	c := Circle{10.5}
	fmt.Println("Call function PerimeterCircle:", PerimeterCircle(c))
	fmt.Println("Call method Perimeter of circle:", c.Perimeter())
	r := Rectangle{10, 20}
	fmt.Println("Call function PerimeterRectangle:", PerimeterRectangle(r))
	fmt.Println("Call method Perimeter of rectangle:", r.Perimeter())
}
