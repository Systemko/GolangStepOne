package main

import "fmt"

func main() {
	// Автоматическая инициализация
	var age int
	fmt.Println("My age is:", age)

	// Ручная инициализация
	var height int = 183
	fmt.Println("My height is:", height)

	// "Полустрогая" типизация
	var uid = 12345
	fmt.Println("My uid:", uid)
	fmt.Printf("Uid type is: %T\n", uid)

	// Типизированные и нетипизированные константы
	const price, tax float32 = 275.00, 27.50
	const quantity, inStock = 2, true
	fmt.Println("Total:", 2*quantity*(price+tax))

	// Изменение значения переменной с использованием литерала
	var cost float32 = 275.00
	var total float32 = 27.50
	cost = 300
	fmt.Println("Cost + Total:", cost+total)

	// Тип с плавающей точкой по умолчанию
	var value = 275.00
	var taxes float32 = 27.50
	fmt.Printf("Value type %T not equal taxes type %T\n", value, taxes)

	// Приведение типов
	var value2 = 275.00
	var taxes2 float32 = 27.50
	fmt.Println("value2 + taxes2:", value2+float64(taxes2))

	// Короткое присваивание
	result := false
	value, new_value := 3.12, 121
	fmt.Println("Results:", value, new_value, result)

	// Ввод данных
	var (
		number int
		s      string
	)
	fmt.Scan(&number)
	fmt.Scan(&s)
	fmt.Println(number, s)
}
