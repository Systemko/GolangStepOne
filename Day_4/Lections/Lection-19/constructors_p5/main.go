package main

import "fmt"

// Создадим тип Rectangle
type Rectangle struct {
	length, width int
}

// Создадим конструктор для Rectangle
// Для имен конструкторов в GO договорились использовать функции
// со следующим названием:
// - New() - если данный конструктор на файл один
// - New<StructName>() - если в файле присутствуют разные структуры
// В GO принято из конструктора возвращать не сам экземпляр, а ссылку на него

func NewRectangle(newLength, newWidth int) *Rectangle {
	return &Rectangle{newLength, newWidth}
}

// Добавим два метода
func (r *Rectangle) Perimeter() int {
	return 2 * (r.length + r.width)
}

func (r *Rectangle) Square() int {
	return r.length * r.width
}

func main() {
	r := NewRectangle(10, 20)

	fmt.Printf("Type: %T and value: %v\n", r, r)
	fmt.Println("Perimeter:", r.Perimeter())
	fmt.Println("Square:", r.Square())
}
