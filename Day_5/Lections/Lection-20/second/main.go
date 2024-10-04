package main

import "fmt"

// А что если создать интерфейс, в котором вообще нет
// никаких требований?

type Empty interface{}

// Кто будет удовлетворять этому типу? Ответ - кто угодно.

// Реализуем функцию, которая может принимать что угодно
func Describer(pretendent interface{}) {
	fmt.Printf("Type = %T and value %v\n", pretendent, pretendent)
}

type Student struct {
	name string
}

// Type Assertion | Утверждение типов
func SemiGeneric(pretendent interface{}) {
	value, ok := pretendent.(int) // Проверка того, что претендент это int
	fmt.Println("Value:", value, "OK?", ok)
}

func main() {
	strSample := "Hello world"

	// Передача параметра без предварительного присваивания в промежуточную переменную
	Describer(strSample)

	intSample := 200
	Describer(intSample)

	Describer(Student{"Vasya"})

	// Работа с полу-дженериком
	SemiGeneric(10)
	SemiGeneric(Student{"Fedor"})
	SemiGeneric("Hello")
}
